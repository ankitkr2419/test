#!/bin/bash
read -p "please input the artifact version: " ver
echo "creting a temporary directory named ${ver}"
mkdir ${ver}
echo "copying migrations,conf,cpagent and installation.sh to ${ver}"
cp -r migrations conf cpagent utils/installation.sh ${ver}
echo "deleting old artifact.zip if any"
rm -f artifact.zip
echo "creating artifact.zip"
zip -r artifact.zip ${ver}
echo "deleting folder ${ver}"
rm -rf ${ver}