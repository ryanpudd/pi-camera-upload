#!/bin/bash

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

cp $SCRIPTPATH/../bin/* /usr/local/bin/
cp $SCRIPTPATH/pi_camera_control /usr/local/bin/

cp $SCRIPTPATH/../config/pi_camera_monitor.service /etc/systemd/system/

systemctl daemon-reload
systemctl restart pi_camera_monitor