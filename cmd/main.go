package main

import (
	api_gateway "homify-go-grpc/internal/api-gateway"
	authentication_service "homify-go-grpc/internal/authentication-service"
	property_service "homify-go-grpc/internal/property-service"
	search_service "homify-go-grpc/internal/search-service"
)

func main() {
	go authentication_service.RunGRPCAuthenticationServer()
	go property_service.RunGRPCPropertyServer()
	go search_service.RunSearchServer()
	api_gateway.RunHTTPServer()
}
