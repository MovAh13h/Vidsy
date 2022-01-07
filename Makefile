.PHONY: proto cmd

all: build

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative ./proto/vidsy.proto

buildw:
	go build ./cmd/vidworker/...

runw: buildw
	clear
	./vidworker -port 3001

buildm:
	go build ./cmd/vidmaster/...

build: proto buildw buildm

clean:
	@rm -rf ./vidworker.exe ./vidworker ./vidmaster.exe ./vidmaster