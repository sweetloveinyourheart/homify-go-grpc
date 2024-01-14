package main

import (
	api_gateway "homify-go-grpc/internal/api-gateway"
	authentication_service "homify-go-grpc/internal/authentication-service"
)

func main() {
	go authentication_service.RunGRPCAuthenticationServer()
	api_gateway.RunHTTPServer()
}
