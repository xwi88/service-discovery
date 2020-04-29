# Makefile to build the command lines and tests in Seele project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.

SHELL := /bin/bash

#BASEDIR = $(shell pwd)
BASEDIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

ldflagsDebug=""

# -s -w
ldflagsRelase="-s -w"

buildTags=""
#buildTags="jsoniter"

.PHONY: proto

defualt: demo-default

all: proto demo-default

#protoc -I proto/helloworld/ proto/helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
#protoc -I proto/ proto/helloworld.proto --go_out=plugins=grpc:helloworld

clean:
	rm -rf pb/go/model/*
	rm -rf pb/go/service/*
	rm -rf pb/java/*

proto:
	protoc -I proto/ --go_out=pb/go/model/ --go_opt=paths=source_relative proto/*.proto
	@echo "Done proto built"

grpc:
	protoc -I proto helloworld.proto --go_out=plugins=grpc:pb/go/service

java:
	protoc -I=proto --java_out=pb/java helloworld.proto

grpc-java:
	protoc -I proto/helloworld helloworld.proto --java_out=plugins=grpc:pb/java/service

clean-demo:
	rm build/bin/demo_*

demo-default:
	mkdir -p build/bin
	go build -v -tags ${buildTags} -ldflags ${ldflagsDebug} -o ./build/bin/demo_client ./demo/client/client.go && \
	go build -v -tags ${buildTags} -ldflags ${ldflagsDebug} -o ./build/bin/demo_server ./demo/server/server.go
	@echo "Done demo built"

demo-darwin:
	mkdir -p build/bin
	export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=amd64 && \
	go build -v -tags ${buildTags} -ldflags ${ldflagsDebug} -o ./build/bin/demo_client_darwin ./demo/client/client.go && \
	go build -v -tags ${buildTags} -ldflags ${ldflagsDebug} -o ./build/bin/demo_server_darwin ./demo/server/server.go
	@echo "Done demo built for darwin"

demo-linux:
	mkdir -p build/bin
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && \
	go build -v -tags ${buildTags} -ldflags ${ldflagsDebug} -o ./build/bin/demo_client_linux ./demo/client/client.go && \
	go build -v -tags ${buildTags} -ldflags ${ldflagsDebug} -o ./build/bin/demo_server_linux ./demo/server/server.go
	@echo "Done demo built for linux"

server1:
	cd demo/server && go run server.go --port 50051

server2:
	cd demo/server && go run server.go --port 50052

server3:
	cd demo/server && go run server.go --port 50053

client:
	cd demo/client && go run client.go

server1-dev:
	cd demo/server && go run server.go --port 50051 -env dev

server2-dev:
	cd demo/server && go run server.go --port 50052 -env dev

server3-dev:
	cd demo/server && go run server.go --port 50053 -env dev

client-dev:
	cd demo/client && go run client.go -env dev