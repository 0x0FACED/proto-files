PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc

.PHONY: run all clean proto build server client

all: proto build

clean:
	rm -rf link_service/gen

proto:
	protoc -I. \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--go_out=link_service/  \
		--go-grpc_out=link_service/  \
		link_service/proto/linkservice.proto