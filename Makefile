.PHONY: proto cmd

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative ./proto/vidsy.proto

buildw:
	go build ./cmd/vidworker/...

buildm:
	go build ./cmd/vidmaster/...

build: proto buildw buildm

clean:
	@rm -rf ./vidworker.exe ./vidworker ./vidmaster.exe ./vidmaster