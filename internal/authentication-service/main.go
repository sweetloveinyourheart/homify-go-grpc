package authentication_service

import (
	"fmt"
	"log"
	"net"

	proto "homify-go-grpc/api/authentication"
	"homify-go-grpc/internal/authentication-service/configs"
	"homify-go-grpc/internal/authentication-service/database"
	"homify-go-grpc/internal/authentication-service/servers"

	"google.golang.org/grpc"
)

func RunGRPCAuthenticationServer() {
	configurations := configs.GetConfig()

	database.InitPostgresConnection()

	lis, err := net.Listen("tcp", configurations.TCPAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := servers.NewGRPCAuthenticationServer()
	proto.RegisterAuthenticationServer(s, srv)

	fmt.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
