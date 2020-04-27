# Makefile to build the command lines and tests in Seele project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.

SHELL := /bin/bash

#BASEDIR = $(shell pwd)
BASEDIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

ldflagsDebug=""

# -s -w
ldflagsRelase="-s -w"

#buildTags=""
buildTags="jsoniter"

.PHONY: proto

defualt: app

all: proto app app-darwin app-linux

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

app:
	go build -v -mod=vendor -tags ${buildTags} -ldflags ${ldflagsDebug} -o ./build/bin/${CMD_APP_NAME} ./cmd/${CMD_APP_DIR}
	@echo "Done app built remain gdb info"


app-darwin:
	export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=amd64 && \
	go build -v -mod=vendor -tags ${buildTags} -ldflags ${ldflagsRelase} -o ./build/bin/${CMD_APP_NAME}-darwin ./cmd/${CMD_APP_DIR}
	@echo "Done app built for darwin"

app-linux:
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && \
	go build -v -mod=vendor -tags ${buildTags} -ldflags ${ldflagsRelase} -o ./build/bin/${CMD_APP_NAME}-linux ./cmd/${CMD_APP_DIR}
	@echo "Done app built for linux"

server1:
	cd demo/server && \
	go run server.go --port 50051

server2:
	cd demo/server && \
	go run server.go --port 50052

server3:
	cd demo/server && \
	go run server.go --port 50053

client:
	cd demo/client && \
    go run client.go