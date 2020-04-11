.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_upload ./cmd/pi_camera_upload
	GOOS=linux CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_monitor ./cmd/pi_camera_monitor

.PHONY: build_arm
build_arm:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_upload ./cmd/pi_camera_upload
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_monitor ./cmd/pi_camera_monitor

.PHONY: install
install:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v
