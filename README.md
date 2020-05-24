# Pi Camera Upload
![Go](https://github.com/ryanpudd/pi-camera-upload/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryanpudd/pi-camera-upload)](https://goreportcard.com/report/github.com/ryanpudd/pi-camera-upload)

A simple application that uploads photos from the Raspberry Pi Camera/Web Cam to an S3 bucket with notifications on slack, it is designed to be triggered by Motion

A seperate monitoring application provides
- Status notifications upon request

![example](./resources/example.png)

- [Prerequisites](#prerequisites)
  - [Raspberry pi](#raspberry-pi)
  - [USB Webcam](#usb-webcam)
  - [Golang](#golang)
- [Device Setup](#device-setup)
  - [Install dependencies](#install-dependencies)
  - [Configure motion](#configure-motion)
- [S3 setup](#s3-setup)
- [Slack setup](#slack-setup)
- [Uploading images](#uploading-images)


# Prerequisites

## Raspberry Pi

[RaspberryPi 3](https://www.raspberrypi.org/products/raspberry-pi-3-model-b/) is used for this build.
They include builtin Wi-Fi, making them well suited for this usecase.

Configure the Raspberry Pi as follows:
- [Raspbian Stretch Lite](https://www.raspberrypi.org/downloads/raspbian/) is the OS this guide was tested with
- Setup ssh access to the pi
- Connect the pi to your wireless network

Setting up ssh access and configuring the wireless is outside of the scope of this guide.
I've found [this guide](https://medium.com/@danidudas/install-raspbian-jessie-lite-and-setup-wi-fi-without-access-to-command-line-or-using-the-network-97f065af722e) to 
be helpful for that.

## Camera

Either the Raspberry Pi Camera or a USB webcam is used for this project. Any USB camera should do, I'm using a [Logitech C920 Webcam](https://www.amazon.com/gp/product/B006JH8T3S).

# Device Setup

TODO

## Golang

You'll need a working Go environment on your local machine if you intend to build the code.

* Go v1.13
* [dep](https://github.com/golang/dep)


## Install dependencies


```
sudo apt-get update

# usb camera support
sudo apt-get install fswebcam

# Install motion
sudo apt-get install motion
```

Test the camera with `fswebcam`
```
fswebcam image.jpg
```
```
--- Opening /dev/video0...
Trying source module v4l2...
/dev/video0 opened.
No input was specified, using the first.
Adjusting resolution from 384x288 to 352x288.
--- Capturing frame...
Captured frame in 0.00 seconds.
--- Processing captured image...
Writing JPEG image to 'image.jpg'.
```
Now we know our webcam is at `/dev/video0`. If you look at `image.jpg` you'll see the picture it took.


## Configure motion

[Motion](https://motion-project.github.io/) is used to monitor the camera.

Edit the following settings in the motion configuration file at `/etc/motion/motion.conf`

```
# Make sure the proper camera device is set
videodevice /dev/video0

# Tell motion to run as a background service
daemon on
```

Edit `/etc/default/motion` and set `start_motion_daemon=yes`
```
start_motion_daemon=yes
```

And restart the system
```
reboot
```

Once it comes back up open another ssh connection.

Confirm motion is running
```
systemctl status motion
```

At this point motion is running as a background service and will be automatically
started after a reboot. The next section will guide you through uploading the images
to AWS S3 and notifying a Slack channel.

# S3 Setup

1. Create an S3 bucket, noting the name and region
2. Create an IAM user, noting the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
3. Create and attach an IAM policy to the user, allowing them to upload to the S3 bucket
  ```
  {
      "Version": "2012-10-17",
      "Statement": [
          {
              "Sid": "cctv_upload",
              "Effect": "Allow",
              "Action": [
                  "s3:PutObject",
                  "s3:PutObjectAcl"
              ],
              "Resource": "arn:aws:s3:::YOUR-BUCKET-NAME/*"
          }
      ]
  }
  ```

# Slack Setup

1. Create a [Legacy Slack Workspace](https://slack.com/intl/en-gb/help/articles/115005265703-Create-a-bot-for-your-workspace)
2. Make note of the bot token

# Uploading images

TODO - Update this guide once deployment mechanism is complete
Create the `env.conf` script with your AWS and Slack secrets:
```
# /usr/local/etc/pi_camera_control/env.conf

SLACK_TOKEN=xxxx
SLACK_CHANNEL_ID=xxxxx
S3_BUCKET_NAME=my-bucket-name
S3_BUCKET_REGION=us-west-2
AWS_ACCESS_KEY=xxx
AWS_SECRET_KEY=xxxxx
```

Build the program that ties the system together.
It's called by the `on_picture_save` script when a new image is created. 
The program uploads the image to S3 and then sends a webhook to Slack with the image url.

```
make install
make build
```

Upload both programs to the pi
```
rsync -avz root/usr/local/bin/ pi@your_ip:/tmp

cp /tmp/cctv_upload /usr/local/bin/
cp /tmp/on_picture_save /usr/local/bin/
```

Finally, open the motion config again and configure it to call the `on_picture_save` script everytime it creates an image.

```
# /etc/motion/motion.conf
on_picture_save /usr/local/bin/pi_camera_control pi_camera_upload %f
```

NOTE: Enable the daemon for the service
NOTE: Output best picture
```
output_pictures best
```

NOTE: Disable Video
```
ffmpeg_output_movies off
```

NOTE: Increase resolution to 1920 x 1080
```
width 1920
height 1080
```

NOTE:
Text needs to be double sized
```
text_double on
```

Restart motion
```
systemctl restart motion
```
