NAME=mc
IMAGE_NAME=micro/$(NAME)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_IMPORT=github.com/justcy/go-micro/cmd
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_IMPORT).BuildDate=$(BUILD_DATE) -X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT) -X $(GIT_IMPORT).GitTag=$(GIT_TAG)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)
PROTO_FLAGS=--go_opt=paths=source_relative --micro_opt=paths=source_relative
PROTO_PATH=$(GOPATH)/src:.
SRC_DIR=$(GOPATH)/src

all: clean build
build:
	go build -a -installsuffix cgo -ldflags "-s -w ${LDFLAGS}" -o $(NAME)
test: vet
	go test -v ./...
.PHONY: clean
clean:
	rm -rf ./mc
