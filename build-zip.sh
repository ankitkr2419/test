read -p "please input the zip version: " ver
echo "creting a temporary directory named ${ver}"
mkdir ${ver}
echo "copying migrations,conf,cpagent and installation.sh to ${ver}"
cp -r migrations conf cpagent installation.sh ${ver}
echo "creating artifact.zip"
zip artifact.zip ${ver}/*
echo "deleting ${ver} folder"
rm -r ${ver}