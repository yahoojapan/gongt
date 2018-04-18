NAME := gongt
VERSION := v0.0.1
GO_VERSION := $(shell go version)
REVISION := $(shell git rev-parse --short HEAD)
PROJECT_ROOT := $(shell pwd)

init:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build:
	go build gongt.go

test:
	go test -v .

bench:
	go run example/main.go -create
	go test -bench .

download:
	cd $(PROJECT_ROOT)/assets/bench && ./download.sh

clean:
	rm -rf $(PROJECT_ROOT)/assets/bench/*.hdf5

.PHONY: build test publish clean bench download init
