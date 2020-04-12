.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_upload ./cmd/pi_camera_upload_x86
	GOOS=linux CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_monitor ./cmd/pi_camera_monitor_x86

.PHONY: build_arm
build_arm:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_upload ./cmd/pi_camera_upload_arm
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_monitor ./cmd/pi_camera_monitor_arm

.PHONY: install
install:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v
