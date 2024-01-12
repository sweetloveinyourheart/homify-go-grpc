PROTO_DIRS := $(wildcard api/*)
PROTO_FILES := $(foreach dir, $(PROTO_DIRS), $(wildcard $(dir)/*.proto))

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --proto_path=api/ $(PROTO_FILES)

.PHONY: proto
