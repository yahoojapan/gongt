NAME := gongt
VERSION := v0.0.1
GO_VERSION := $(shell go version)
REVISION := $(shell git rev-parse --short HEAD)
PROJECT_ROOT := $(shell pwd)

build:
	go build gongt.go

test:
	go test -v .

bench:
	cd $(PROJECT_ROOT)/assets/bench && ./download.sh
	go run example/main.go -create
	go test -bench .

clean:

.PHONY: build test publish clean bench
