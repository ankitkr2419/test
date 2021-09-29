#!/bin/bash
read -p "please input the artifact version: " ver
echo "creting a temporary directory named ${ver}"
mkdir -p ${ver}/utils
echo "copying migrations,conf,cpagent and installation.sh to ${ver}"
cp -r migrations conf cpagent utils/installation.sh  utils/releases.txt ${ver}
echo "copying covid recipes to ${ver}"
cp utils/Covid_Ext_v1.4.0.csv utils/Covid_PCR_v1.4.0.csv ${ver}/utils
echo "deleting old artifact.zip if any"
rm -f artifact.zip
echo "creating artifact.zip"
zip -r artifact.zip ${ver}
echo "deleting folder ${ver}"
rm -rf ${ver}