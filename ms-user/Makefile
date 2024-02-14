# Makefile

# Path to protobuf compiler
PROTOC := protoc

# Proto files directory
PROTO_DIR := proto

# Output directory
OUTPUT_DIR := pb

# Proto files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Generate Go files
generate:
    $(PROTOC) --go_out=$(OUTPUT_DIR) --go-grpc_out=$(OUTPUT_DIR) $(PROTO_FILES)

.PHONY: generate
