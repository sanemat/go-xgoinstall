VERSION = $(shell gobump show -r)
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X github.com/sanemat/go-xgoinstall.revision=$(CURRENT_REVISION)"
u := $(if $(update),-u)

.PHONY: test
test: download
	go test

.PHONY: download
download:
	echo Download go.mod dependencies
	go mod download
	go mod tidy

.PHONY: install-tools
install-tools: download
	echo Installing tools from tools.go
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: goimports
goimports:
	goimports -w .

echo:
	echo ${VERSION} ${BUILD_LDFLAGS}

.PHONY: build
build: download
	go build -ldflags=$(BUILD_LDFLAGS) ./cmd/x-go-install

.PHONY: install
install: download
	go install -ldflags=$(BUILD_LDFLAGS) ./cmd/x-go-install

.PHONY: crossbuild
crossbuild:
	goxz -pv=v$(VERSION) -build-ldflags=$(BUILD_LDFLAGS) \
      -os=linux,darwin,windows -d=./dist/v$(VERSION) ./cmd/*

.PHONY: upload
upload:
	ghr v$(VERSION) dist/v$(VERSION)

.PHONY: credits
credits:
	gocredits . > credits.txt
