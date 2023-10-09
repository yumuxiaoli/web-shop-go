
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: update
update:
	@go get -u

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/cart.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/cart.proto
	
.PHONY: build
build:
	go build -o cart *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cart:latest
