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
dir=${HOME}"/cpagent/utils/logs/"
var=${HOME}"/cpagent/utils/logs/output_$(date +%s).log"

create_log_file () {
        echo "Logs for current run are present in "${var}
        mkdir -p ${dir}
        touch ${var}
}

if [[ ${apptype} == "extraction" ]]; then
    create_log_file
    sudo ./cpagent start --plc ${dev} --no-rtpcr ${delay} > ${var}
elif [[ ${apptype} == "rtpcr" ]]; then
    create_log_file
    sudo ./cpagent start --plc ${dev} --tec ${dev} --no-extraction > ${var}
else
  echo "Please specify apptype[extraction, rtpcr]"
  exit 1
fi

echo "Logs for current run are present in "${var}
