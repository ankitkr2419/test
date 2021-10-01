#!/bin/bash
export apptype=${1}
export dev=${2}
export delay=" --delay 10 "

# Delay is not allowed for compact
# By default compact32 will start
if [[ ${dev} == "" ]]; then
    dev="compact32"
    delay=""
elif [[  ${dev} != "simulator" ]]; then
    echo "Please specify plc type as simulator"
    exit 1
fi

chmod +x ./cpagent
var="./utils/logs/output_$(date +%s).log"

if [[ ${apptype} == "extraction" ]]; then
    echo "Logs for current run are present in "${var}
    sudo ./cpagent start --plc ${dev} --no-rtpcr ${delay} > ${var}
elif [[ ${apptype} == "rtpcr" ]]; then
    echo "Logs for current run are present in "${var}
    sudo ./cpagent start --plc ${dev} --tec ${dev} --no-extraction > ${var}
else
  echo "Please specify apptype[extraction, rtpcr]"
  exit 1
fi

echo "Logs for current run are present in "${var}