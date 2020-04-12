.PHONY: release install test build_all build build_arm $(PLATFORMS)
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_upload ./cmd/pi_camera_upload
	GOOS=linux CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_monitor ./cmd/pi_camera_monitor

build_arm:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_upload ./cmd/pi_camera_upload
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./root/usr/local/bin/pi_camera_monitor ./cmd/pi_camera_monitor

PLATFORMS := linux/amd64 linux/arm windows/amd64
OUT_DIR = bin


temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

$(PLATFORMS): init
	mkdir $(OUT_DIR)/$(os)-$(arch)
	GOOS=$(os) GOARCH=$(arch) go build -o '$(OUT_DIR)/$(os)-$(arch)' ./cmd/...

init:
	mkdir -p $(OUT_DIR)
clean:
	rm -rf $(OUT_DIR)
build_all: $(PLATFORMS)
release: test build_all
test:
	CGO_ENABLED=0 go test ./cmd/...
zip:
	zip $(OUT_DIR)

install:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v
