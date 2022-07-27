RELEASE_VER ?= latest
BUILD_DATE  := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
BASE_DIR    := $(shell git rev-parse --show-toplevel)
GIT_SHA     := $(shell git rev-parse --short HEAD)
BIN         := $(BASE_DIR)/bin

ifndef PKGS
PKGS := $(shell go list ./... 2>&1 | grep -v 'github.com/portworx/kdmp/vendor' | grep -v versioned | grep -v 'pkg/apis/v1')
endif

GO_FILES := $(shell find . -name '*.go' | grep -v 'vendor' | \
                                   grep -v '\.pb\.go' | \
                                   grep -v '\.pb\.gw\.go' | \
                                   grep -v 'externalversions' | \
                                   grep -v 'versioned' | \
                                   grep -v 'generated')

.DEFAULT_GOAL: all
.PHONY: test
all: pretest test

test: unittest

build:
	@echo "Build nfs-executor"
	go build -o ${BIN}/nfs-executor -ldflags="-s -w \
	-X github.com/portworx/kdmp/pkg/version.gitVersion=${RELEASE_VER} \
	-X github.com/portworx/kdmp/pkg/version.gitCommit=${GIT_SHA} \
	-X github.com/portworx/kdmp/pkg/version.buildDate=${BUILD_DATE} \
	-X github.com/portworx/kdmp/pkg/version.major=${MAJOR_VERSION} \
	-X github.com/portworx/kdmp/pkg/version.minor=${MINOR_VERSION} \
	-X github.com/portworx/kdmp/pkg/version.patch=${PATCH_VERSION}" \
	-a $(BASE_DIR)/cmd/executor/nfs


staticcheck:
	go get -u honnef.co/go/tools/cmd/staticcheck
	staticcheck $(PKGS)
	staticcheck -tags integrationtest test/integration_test/*.go
	staticcheck -tags unittest $(PKGS)


errcheck:
	go get -u github.com/kisielk/errcheck
	errcheck -ignoregenerated -verbose -blank $(PKGS) 
	errcheck -ignoregenerated -verbose -blank -tags unittest $(PKGS)
	errcheck -ignoregenerated -verbose -blank -tags integrationtest github.com/portworx/px-backup/test/integration_test


check-fmt:
	bash -c "diff -u <(echo -n) <(gofmt -l -d -s -e $(GO_FILES))" 


vet:
	go vet $(PKGS)
	go vet -tags unittest $(PKGS)
	go vet -tags integrationtest github.com/portworx/px-backup/test/integration_test


do-fmt:
	gofmt -s -w $(GO_FILES)

pretest: check-fmt lint vet errcheck staticcheck