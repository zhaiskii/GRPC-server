PROTO_DIR=.
OUT_DIR=pkg/api/test/api
PROTO_FILES=$(PROTO_DIR)/*.proto

BUILD_DIR=bin
BINARY=$(BUILD_DIR)/server

GOBIN=$(shell go env GOPATH)/bin
GRPC_GATEWAY_GEN=$(GOBIN)/protoc-gen-grpc-gateway
PROTOC_GEN_GO=$(GOBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC=$(GOBIN)/protoc-gen-go-grpc

.PHONY: all proto build run clean

all: proto build
proto:
	@echo "Generating gRPC and gRPC Gateway files..."
	protoc -I $(PROTO_DIR) -Igoogleapis \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(OUT_DIR) --grpc-gateway_opt=paths=source_relative \
		$(PROTO_FILES)

build: proto
	@echo "Building binary..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BINARY) ./cmd/server

run: build
	@echo "Running server..."
	$(BINARY)

