#!/bin/bash -xv
export action=${1}
export curr_ver=`pwd`
if [[ -f $HOME/cpagent/old_version.txt ]]; then
        export old_ver=$(cat $HOME/cpagent/old_version.txt| head -1)
fi
upgrade() {
        echo "This is upgrade function  :"
        echo "stoping cpagent service "
        echo "admin" | sudo -s systemctl stop cpagent
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
                echo "Symlink is present"
                rm $HOME/cpagent/current
                ln -sf $curr_ver $HOME/cpagent/current
        else
                echo "Symlink current not exists, creating symlink with name current"
                ln -sf $curr_ver $HOME/cpagent/current
        fi
        ### Run migrations
        cd $HOME/cpagent/current
        ./cpagent migrate & pid=$!
        wait $pid
        #### Run go binary or start cpagent service
        #echo "admin" | nohup sudo -S ./cpagent start --plc simulator &#
        echo "admin" | sudo systemctl start cpagent
        #kill $(ps aux  |grep gedit |grep -v grep |awk {'print $2'})
}
revert() {
        echo "This is revert function :"
        echo "stop running process"
        if [[ -L $HOME/cpagent/current ]]; then
                echo "Symlink is present"
                rm $HOME/cpagent/current

                ln -sf ${old_ver} $HOME/cpagent/current
        else
                echo "Symlink current not exists, creating with name current"
                ln -sf ${old_ver} $HOME/cpagent/current
        fi
}
##### Main #####

if [[ ! -d $HOME/cpagent/release ]]; then
        mkdir $HOME/cpagent/release
fi
if [[ ! -d $HOME/logs ]]; then
        mkdir -p $HOME/logs
fi
if [[ ! -d $HOME/logs/$log_dir ]]; then
        mkdir -p $HOME/logs/$log_dir
fi
#if [[ ! -d mylab/status ]]; then
#        mkdir -p $HOME/status
#fi
exec > >(tee "mylab/logs/${log_dir}/installation_$(date +"%m%d%Y-%T").log") 2>&1

if [[ ${action} == "upgrade" ]]; then

  upgrade

elif [[ ${action} == "revert" ]]; then

  revert

else

  echo "No action specified "
  exit 1
fi
echo "Installation script is completed" 1>&2
