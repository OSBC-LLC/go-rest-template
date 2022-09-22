#!/bin/bash

read -p "Enter module name: " VALUE
echo "replacing 'github.com/SFDC/orch-rest-template' with:"
echo "           $VALUE"

rg github.com/SFDC/orch-rest-template -l -g '!init.sh' | xargs sed -i -e "s+github.com/SFDC/orch-rest-template+$VALUE+g"

rm **/**-e
