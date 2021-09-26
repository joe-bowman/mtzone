TEST_DOCKER_REPO=microtick/mtzonetest

export GO111MODULE = on
GO := go

UNAME_OS              := $(shell uname -s)
UNAME_ARCH            := $(shell uname -m)
CACHE_BASE            ?= $(abspath .cache)
CACHE                 := $(CACHE_BASE)
CACHE_BIN             := $(CACHE)/bin
CACHE_INCLUDE         := $(CACHE)/include
CACHE_VERSIONS        := $(CACHE)/versions

BUF_VERSION           ?= 0.20.5
PROTOC_VERSION        ?= 3.11.2

# <TOOL>_VERSION_FILE points to the marker file for the installed version.
# If <TOOL>_VERSION_FILE is changed, the binary will be re-downloaded.
BUF_VERSION_FILE           = $(CACHE_VERSIONS)/buf/$(BUF_VERSION)
PROTOC_VERSION_FILE        = $(CACHE_VERSIONS)/protoc/$(PROTOC_VERSION)

MODVENDOR = $(CACHE_BIN)/modvendor
BUF := $(CACHE_BIN)/buf
PROTOC := $(CACHE_BIN)/protoc

$(eval override VERSION = $(shell git describe --tags --always 2>/dev/null))
$(eval override DATE = $(shell date))
$(eval override HOST = $(shell hostname))
$(eval override COMMIT = $(shell git log -1 --format='%H'))
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=Microtick \
	  -X github.com/cosmos/cosmos-sdk/version.AppName=mtm \
	  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=build_host=$(HOST);build_date=$(DATE)"


all: install

install:
	$(GO) install -mod=mod -tags="netgo ledger" -ldflags='$(ldflags)' ./cmd/mtm
	@mv $(shell go env GOPATH)/bin/mtm .

ifeq ($(UNAME_OS),Linux)
  PROTOC_ZIP ?= protoc-${PROTOC_VERSION}-linux-x86_64.zip
endif
ifeq ($(UNAME_OS),Darwin)
  PROTOC_ZIP ?= protoc-${PROTOC_VERSION}-osx-x86_64.zip
endif
	
.PHONY: proto
proto: $(PROTOC) protovendor
	@echo "Installing protoc-gen-gocosmos..."
	go mod tidy
	@go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos
	@echo "Installing protoc-gen-grpc-gateway"
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@echo "Creating protobuf classes"
	@PATH="$(PATH)":$(shell go env GOPATH)/bin ./scripts/protocgen.sh

.PHONY: js
js: $(PROTOC) protovendor
	@echo "Installing protoc-gen-gocosmos..."
	@go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos
	@echo "Installing protoc-gen-grpc-gateway"
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@echo "Creating javascript archive"
	@PATH=$(PATH):$(shell go env GOPATH)/bin ./scripts/protocjs.sh
	
.PHONY: protovendor
protovendor: modsensure $(MODVENDOR)
	@echo "vendoring *.proto files..."
	$(MODVENDOR) -copy="**/*.proto" -include=\
github.com/cosmos/cosmos-sdk/proto,\
github.com/cosmos/cosmos-sdk/third_party/proto,\
github.com/tendermint/tendermint/proto,\
github.com/gogo/protobuf,\
github.com/regen-network/cosmos-proto/cosmos.proto

$(CACHE):
	@echo "creating .cache dir structure..."
	mkdir -p $@
	mkdir -p $(CACHE_BIN)
	mkdir -p $(CACHE_INCLUDE)
	mkdir -p $(CACHE_VERSIONS)

.PHONY: modsensure
modsensure: deps-tidy deps-vendor

deps-tidy:
	$(GO) mod tidy

deps-vendor:
	$(GO) mod vendor
	
$(BUF_VERSION_FILE): $(CACHE)
	@echo "installing protoc buf cli..."
	rm -f $(BUF)
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_OS)-$(UNAME_ARCH)" \
		-o "$(CACHE_BIN)/buf"
	chmod +x "$(CACHE_BIN)/buf"
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(BUF): $(BUF_VERSION_FILE)

$(PROTOC_VERSION_FILE): $(CACHE)
	@echo "installing protoc compiler..."
	rm -f $(PROTOC)
	(cd /tmp; \
	curl -sOL "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"; \
	unzip -oq ${PROTOC_ZIP} -d $(CACHE) bin/protoc; \
	unzip -oq ${PROTOC_ZIP} -d $(CACHE) 'include/*'; \
	rm -f ${PROTOC_ZIP})
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC): $(PROTOC_VERSION_FILE)
	
$(MODVENDOR): $(CACHE)
	echo "installing modvendor..."
	GOBIN=$(CACHE_BIN) go get github.com/goware/modvendor

test-docker:
	@docker build -f contrib/Dockerfile.test -t ${TEST_DOCKER_REPO}:$(shell git rev-parse --short HEAD) .
	@docker tag ${TEST_DOCKER_REPO}:$(shell git rev-parse --short HEAD) ${TEST_DOCKER_REPO}:$(shell git rev-parse --abbrev-ref HEAD | sed 's#/#_#g')
	@docker tag ${TEST_DOCKER_REPO}:$(shell git rev-parse --short HEAD) ${TEST_DOCKER_REPO}:latest
	@docker push ${TEST_DOCKER_REPO}:$(shell git rev-parse --short HEAD)
	@docker push ${TEST_DOCKER_REPO}:$(shell git rev-parse --abbrev-ref HEAD | sed 's#/#_#g')
	@docker push ${TEST_DOCKER_REPO}:latest

clean:
	rm -rf .cache vendor js
