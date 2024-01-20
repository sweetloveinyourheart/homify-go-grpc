package main

import (
	api_gateway "homify-go-grpc/internal/api-gateway"
	authentication_service "homify-go-grpc/internal/authentication-service"
	property_listing_service "homify-go-grpc/internal/property-listing-service"
)

func main() {
	go authentication_service.RunGRPCAuthenticationServer()
	go property_listing_service.RunGRPCPropertyListingServer()
	api_gateway.RunHTTPServer()
}
