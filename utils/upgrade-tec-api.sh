#!/bin/bash -xv

# set current and temp path
current='utils/MeComAPIs/current'
temp='utils/MeComAPIs/temp'

mkdir ${current} ${temp}

# copy ComPort files to MeComAPI/temp
cp utils/MeComAPIs/ComPort/* ${temp}

# copy ConsoleIO files to MeComAPI/temp
cp utils/MeComAPIs/ConsoleIO/* ${temp}

# copy TEC files to MeComAPI/temp
cp utils/MeComAPIs/TEC/* ${temp}

# Let user set zip path
# default path provided
read -e -p "please input the MeComAPI zip's Path: " -i "utils/MeComAPIs/MeComAPI_v0.42/MeComAPI_v0.42.zip" zippath

# unzip the MeComAPI
unzip  ${zippath} -d ${current}

# remove version specific directory names
parent=$(ls ${current})

cp -r ${current}/${parent}/* ${current}

cp ${current}/MeComAPI/private/* ${current}/MeComAPI

# Delete MePort_Winc.c
rm ${current}/MeComAPI/MePort_Win.c

# Rename MePort_Linux.c to just MePort.c
mv ${current}/MeComAPI/MePort_Linux.c ${current}/MeComAPI/MePort.c

# copy MeComAPI folder/subfolders files to MeComAPI/temp
find ${current}/MeComAPI/ -type f \( -name "*.c" -o -name "*.h" \) -exec cp '{}' "utils/MeComAPIs/temp/" ";"

# change include path of Frame.h to current path
# NOTE: replacing early with sed didn't work. Changes didn't reflect in copied files
##TODO: Take suggestion from @krushna to see if there is a better way
sed -i 's/\.\.\/MePort\.h/MePort\.h/g' ${temp}/MeFrame.h
sed -i 's/\.\.\/MePort\.h/MePort\.h/g' ${temp}/MeFrame.c
sed -i 's/private\/MeFrame\.h/MeFrame\.h/g' ${temp}/MePort.c
sed -i 's/\.\.\/ComPort\/ComPort.h/ComPort\.h/g' ${temp}/MePort.c
sed -i 's/\.\.\/MeCom\.h/MeCom\.h/g' ${temp}/MeCom.c

# change include path properly to include all the files under single directory i.e under temp
# Reason: go considers all of C files as a single package
# If there is to be complex directory structure then please use Makefile to manipulate cgo build
# As of now we haven't researched into that land.
# TODO: #include update

# TODO: implement below logic as well via script
# Change `define DEVICE "/dev/ttyUSB" ` in ComPort.c to your needed port interface.
# NOTE: for Linux the ports will be like /dev/ttyUSB0, this means ComPortNr value is 0
# For Windows the ports will be like /dev/S4, this means ComPortNr value is 4

# copy the files in MeComAPI/temp folder to tec_1089. Overwriting if needed.
cp -rf ${temp}/* tec/tec_1089/

# Clear current and temp directory
rm -r ${current} ${temp}

# If any old files are causing trouble while building, then consider deleting/modifying them.

# NOTES
## This script is only guranteed to work with v0.42. For other versions modifications may be needed.
## Tec.c and Tec.h are created by @joshsoftware.com for running the TEC Application.
## These files serve as a medium between go and TEC APIs which are coded in C.
## ComPort and ConsoleIO files are copied from Demo Application which is open-source.

