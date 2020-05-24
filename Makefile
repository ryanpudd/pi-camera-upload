.PHONY: release install test build_all $(PLATFORMS)

PACKAGE_NAME=pi-camera-upload
OUT_DIR = ./bin
DIST_DIR = ./dist
CMD_DIR = ./cmd
SCRIPT_DIR = ./scripts
CONFIG_DIR = ./configs
PLATFORMS := linux/amd64 linux/arm windows/amd64

BUILDS := $(patsubst %, build_%, $(PLATFORMS))
ZIPS := $(patsubst %, zip_%, $(PLATFORMS))
INITS := $(patsubst %, init_%, $(PLATFORMS))
PACKAGES := $(patsubst %, package_%, $(PLATFORMS))

temp = $(subst /, ,$@)
os = $(word 1, $(word 2, $(subst _, , $(temp))))
arch = $(word 2, $(temp))

$(INITS): init
	mkdir $(OUT_DIR)/$(os)-$(arch)

$(BUILDS): $(INITS)
	GOOS=$(os) GOARCH=$(arch) go build -o '$(OUT_DIR)/$(os)-$(arch)' $(CMD_DIR)/...

$(ZIPS): dist
	zip -o $(DIST_DIR)/$(os)-$(arch).zip $(OUT_DIR)/$(os)-$(arch) $(SCRIPT_DIR) $(CONFIG_DIR)

init:
	mkdir -p $(OUT_DIR)
dist:
	mkdir = $(DIST_DIR)
clean:
	rm -rf $(OUT_DIR)
test:
	CGO_ENABLED=0 go test $(CMD_DIR)/...
install:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v

zip: $(ZIPS)
build: $(BUILDS)
release: test build zip
