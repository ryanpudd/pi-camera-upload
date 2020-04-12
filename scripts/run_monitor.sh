# Export the vars in .env into your shell:
export $(egrep -v '^#' .env | xargs)

./root/usr/local/bin/pi_camera_monitor
