package property_listing_service

import (
	"fmt"
	proto "homify-go-grpc/api/property-listing"
	"homify-go-grpc/internal/property-listing-service/configs"
	"homify-go-grpc/internal/property-listing-service/database"
	"homify-go-grpc/internal/property-listing-service/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGRPCPropertyListingServer() {
	configurations := configs.GetConfig()

	db := database.InitPostgresConnection()

	lis, err := net.Listen("tcp", configurations.TCPAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.NewGRPCPropertyListingServer(db)
	proto.RegisterPropertyListingServer(s, srv)

	fmt.Printf("🚀 Property Listing Server is listening on port %s ... \n", configurations.TCPAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
