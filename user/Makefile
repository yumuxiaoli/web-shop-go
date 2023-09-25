
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install micro.dev/v4/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	protoc  --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --micro_out=. --go_o
ut=. --proto_path=. proto/user/*.proto
	
.PHONY: build
build:
	go build -o user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user:latest
