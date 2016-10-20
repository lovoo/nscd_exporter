BUILD_DIR    = $(CURDIR)/build
PROJECT_NAME = nscd_exporter
VERSION      = $(shell git describe --tags || echo 0.0.0-dev)
GO           = go
GOX          = gox
GOX_ARGS     = -output="$(BUILD_DIR)/{{.Dir}}_{{.OS}}_{{.Arch}}" -osarch="linux/amd64 linux/386 linux/arm linux/arm64"

.PHONY: build
build:
	GOBIN=$(BUILD_DIR) $(GO) install -ldflags '-X main.Version=$(VERSION)'

.PHONY: deb
deb:
	make build-deb ARCH=amd64
	make build-deb ARCH=386

.PHONY: build-deb
build-deb:
	fpm -s dir -t deb \
		--name $(PROJECT_NAME) \
		--version $(VERSION) \
		--package $(BUILD_DIR)/$(PROJECT_NAME)_$(VERSION)_$(ARCH).deb \
		--depends nscd \
		--maintainer "LOVOO IT Operations <it.operations@lovoo.com>" \
		--deb-priority optional \
		--category monitoring \
		--force \
 		--deb-compression bzip2 \
		--license "BSD-3-Clause" \
		--vendor "LOVOO GmbH" \
		--deb-no-default-config-files \
		--after-install packaging/postinst.deb \
		--before-remove packaging/prerm.deb \
		--url https://github.com/lovoo/nscd_exporter \
		--description "Exports statistics from NSCD (Name service caching daemon) and publishes them for scraping by Prometheus." \
		--architecture $(ARCH) \
		$(BUILD_DIR)/$(PROJECT_NAME)_linux_$(ARCH)=/usr/bin/ncsd_exporter \
		packaging/nscd-exporter.service=/lib/systemd/system/nscd-exporter.service

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/* || true

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: release-build
release-build:
	$(GO) get -u github.com/mitchellh/gox
	@$(GOX) $(GOX_ARGS) github.com/lovoo/nscd_exporter
