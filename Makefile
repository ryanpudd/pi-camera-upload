PLATFORMS := linux/amd64 linux/arm windows/amd64
.PHONY: release install test build_all $(PLATFORMS)

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
