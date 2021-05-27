#!/bin/bash -xv
export action=${1}
export curr_ver=`pwd`
if [[ -f $HOME/cpagent/old_version.txt ]]; then
        export old_ver=$(cat $HOME/cpagent/old_version.txt| head -1)
fi

safe_to_upgrade () {

while [ "$state" -ne 200 ]
do
        echo -e "\n\n\t: Not safe to upgrade, Run in progress, waiting 5 sec. :"
        sleep 5
        
	export  state=$(curl -o /dev/null -s -w "%{http_code}\n" --location --request GET '0.0.0.0:33001/safe-to-upgrade' --header 'Accept: application/vnd.MyLabDiscoveries.v1')

done

echo -e "\n\n\t: Safe to upgrade :"
echo -e "\n\n\t: Progressing .. : "
}


upgrade() {
        echo -e "\n\t: Started with instllation procedure :"
        echo -e "\n\t: stoping cpagent service :"
        echo "admin" | sudo -S systemctl stop cpagent
        if [[ -f $HOME/cpagent/old_version2.txt ]]; then
                rm $HOME/cpagent/old_version2.txt
        fi
        if [[ -f $HOME/cpagent/old_version.txt ]]; then
                mv $HOME/cpagent/old_version.txt $HOME/cpagent/old_version2.txt
        fi
        if [[ -f $HOME/cpagent/new_version.txt ]]; then
                mv $HOME/cpagent/new_version.txt $HOME/cpagent/old_version.txt
        fi
        echo "$curr_ver" > $HOME/cpagent/new_version.txt
        if [[ -L $HOME/cpagent/current ]]; then
                echo -e "\n\t: Symlink is present :"
                rm $HOME/cpagent/current
                ln -sf $curr_ver $HOME/cpagent/current
        else
                echo -e "\n\t Symlink current not exists, creating symlink with name current :"
                ln -sf $curr_ver $HOME/cpagent/current
        fi
        ### Run migrations
        cd $HOME/cpagent/current
	echo -e "\n\t Running migrations :"
        ./cpagent migrate & pid=$!
        wait $pid
        #### Run go binary or start cpagent service
        #echo "admin" | nohup sudo -S ./cpagent start --plc simulator &#
	echo -e "\n\t Starting cpagent service : "
        echo "admin" | sudo -S systemctl start cpagent
        #kill $(ps aux  |grep gedit |grep -v grep |awk {'print $2'})
}
revert() {
        echo -e "\n\t Started with revert procedure :"
        echo -e "\n\t stoping cpagent service :"
        
        echo "admin" | sudo -S systemctl stop cpagent
        if [[ -L $HOME/cpagent/current ]]; then
                echo -e "\n\t Symlink is present :"
                rm $HOME/cpagent/current

                ln -sf ${old_ver} $HOME/cpagent/current
        else
                echo -e "\n\t Symlink current not exists, creating symlink with name current :"
                ln -sf ${old_ver} $HOME/cpagent/current
        fi
	### Run migrations
        cd $HOME/cpagent/current
	echo -e "\n\t Running migrations :"
        ./cpagent migrate & pid=$!
        wait $pid
        #### Run go binary or start cpagent service
        #echo "admin" | nohup sudo -S ./cpagent start --plc simulator &#
	echo -e "\n\t Starting cpagent service : "
        echo "admin" | sudo -S systemctl start cpagent
        #kill $(ps aux  |grep gedit |grep -v grep |awk {'print $2'})

}
##### Main #####

if [[ ! -d $HOME/cpagent/release ]]; then
        mkdir $HOME/cpagent/release
fi
if [[ ! -d $HOME/cpagent/logs ]]; then
        mkdir -p $HOME/cpagent/logs
fi
if [[ ! -d $HOME/logs/$log_dir ]]; then
        mkdir -p $HOME/logs/$log_dir
fi

exec > >(tee "$HOME/cpagent/logs/${log_dir}/installation_$(date +"%m%d%Y-%T").log") 2>&1

export state=$(curl -o /dev/null -s -w "%{http_code}\n" --location --request GET '0.0.0.0:33001/safe-to-upgrade' --header 'Accept: application/vnd.MyLabDiscoveries.v1')

if [[ $state -eq 417 ]]; then
        echo -e "\n\n\t: Run is in progress :"
        safe_to_upgrade
fi

if [[ $state -eq 000 ]]; then
        echo -e "\n\n\t: Cpagent is in stop state : "
        echo -e "\n\n\t: Progressing .. : "
fi


if [[ ${action} == "upgrade" ]]; then

  upgrade

elif [[ ${action} == "revert" ]]; then

  revert

else

  echo "No action specified "
  exit 1
fi
echo "Installation script is completed" 1>&2
