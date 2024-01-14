.PHONY: all build run clean proto

PROTO_DIRS := $(wildcard api/*)
PROTO_FILES := $(foreach dir, $(PROTO_DIRS), $(wildcard $(dir)/*.proto))

all: build run

build: 
	proto
	echo "Building application..."
	cd cmd && go build -o ../bin

run:
	echo "Running application..."
	./bin &

clean:
	echo "Cleaning up..."
	rm -f bin

swag:
	swag init -g internal/api-gateway/api_gateway.go

proto:
	protoc --go_out=api --go_opt=paths=source_relative \
    --go-grpc_out=api --go-grpc_opt=paths=source_relative \
    --proto_path=api/ $(PROTO_FILES)
