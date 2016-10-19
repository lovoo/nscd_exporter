BUILD_DIR    = $(CURDIR)/build
PROJECT_NAME = nscd_exporter

GO       = go
GOX      = gox
GOX_ARGS = -output="$(BUILD_DIR)/{{.Dir}}_{{.OS}}_{{.Arch}}" -osarch="linux/amd64 linux/386 linux/arm linux/arm64 darwin/amd64 freebsd/amd64 freebsd/386 windows/386 windows/amd64"

.PHONY: build
build:
	GOBIN=$(BUILD_DIR) $(GO) install

.PHONY: clean
clean:
	rm -R $(BUILD_DIR)/* || true

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: release-build
release-build:
	@go get -u github.com/mitchellh/gox
