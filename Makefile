NAME := gongt
VERSION := v0.0.2
GO_VERSION := $(shell go version)
REVISION := $(shell git rev-parse --short HEAD)
PROJECT_ROOT := $(shell pwd)

deps:
	curl -LO https://github.com/yahoojapan/NGT/archive/v$(NGT_VERSION).tar.gz
	tar zxf v$(NGT_VERSION).tar.gz -C /tmp
	cd /tmp/NGT-$(NGT_VERSION); cmake .
	make -j -C /tmp/NGT-$(NGT_VERSION)
	make install -C /tmp/NGT-$(NGT_VERSION)

build:
	go build gongt.go

test:
	go test -v .

bench:
	go test -run ^a -bench BenchmarkFashionMNIST
	go test -run ^a -bench BenchmarkGlove25
	go test -run ^a -bench BenchmarkGlove50
	go test -run ^a -bench BenchmarkGlove100
	#go test -run ^a -bench BenchmarkGlove200
	go test -run ^a -bench BenchmarkNYTimes
	go test -run ^a -bench BenchmarkSIFT

download:
	cd $(PROJECT_ROOT)/assets/bench && ./download.sh

index:
	cd $(PROJECT_ROOT)/example && go build -o $(PROJECT_ROOT)/assets/bench/example
	cd $(PROJECT_ROOT)/assets/bench && ./mkindex.sh

clean:
	rm -rf $(PROJECT_ROOT)/assets/bench/*.hdf5

.PHONY: build test publish clean bench download init
