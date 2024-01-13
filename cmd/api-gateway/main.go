package main

import (
	_ "homify-go-grpc/docs"
	api_gateway "homify-go-grpc/internal/api-gateway"
)

func main() {
	api_gateway.RunHTTPServer()
}
