#!/bin/bash	
chmod +x ./cpagent
var="./utils/logs/output_$(date +%s).log"
sudo ./cpagent start --plc simulator --no-rtpcr --delay 1 > ${var}