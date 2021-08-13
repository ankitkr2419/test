#!/bin/bash

# copy ComPort files to MeComAPI/temp
cp utils/MeComAPIs/ComPort/* utils/MeComAPIs/temp

# copy ConsoleIO files to MeComAPI/temp
cp utils/MeComAPIs/ConsoleIO/* utils/MeComAPIs/temp

# copy TEC files to MeComAPI/temp
cp utils/MeComAPIs/TEC/* utils/MeComAPIs/temp

# unzip the MeComAPI
# TODO: Let user set zip path
unzip -d utils/MeComAPIs/current/ utils/MeComAPIs/MeComAPI_v0.42/MeComAPI_v0.42.zip

# Delete MePort_Winc.c
# TODO: remove version specific directory names
rm utils/MeComAPIs/current/MeComAPI_v0.42/MeComAPI/MePort_Win.c

# Rename MePort_Linux.c to just MePort.c
mv utils/MeComAPIs/current/MeComAPI_v0.42/MeComAPI/MePort_Linux.c utils/MeComAPIs/current/MeComAPI_v0.42/MeComAPI/MePort.c

# copy MeComAPI folder/subfolders files to MeComAPI/temp
find utils/MeComAPIs/current/MeComAPI_v0.42/MeComAPI/ -name "*.c" -o -name "*.h" -exec cp  '{}' "utils/MeComAPIs/temp/" ";"

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
cp -rf utils/MeComAPIs/temp/* tec/tec_1089/

# Clear current and temp directory
rm -r utils/MeComAPIs/current utils/MeComAPIs/temp

# If any old files are causing trouble while building, then consider deleting/modifying them.

# NOTES
## This script is only guranteed to work with v0.42. For other versions modifications may be needed.
## Tec.c and Tec.h are created by @joshsoftware.com for running the TEC Application.
## These files serve as a medium between go and TEC APIs which are coded in C.
## ComPort and ConsoleIO files are copied from Demo Application which is open-source.

