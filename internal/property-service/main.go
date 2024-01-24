package property_service

import (
	"fmt"
	proto "homify-go-grpc/api/property"
	"homify-go-grpc/internal/property-service/configs"
	"homify-go-grpc/internal/property-service/database"
	"homify-go-grpc/internal/property-service/producers"
	"homify-go-grpc/internal/property-service/server"
	kafka_configs "homify-go-grpc/internal/shared/kafka-configs"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGRPCPropertyServer() {
	configurations := configs.GetConfig()
	kafkaConfigs := kafka_configs.GetConfig()

	// Init database
	db := database.InitPostgresConnection()

	// Init kafka producer
	p := producers.NewPropertyProducer(kafkaConfigs)
	p.InitDeliveryReport()
	defer p.CloseProducer()

	// Init property listing service
	lis, err := net.Listen("tcp", configurations.TCPAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.NewGRPCPropertyServer(db, p)
	proto.RegisterPropertyServer(s, srv)

	fmt.Printf("🚀 Property  Server is listening on port %s ... \n", configurations.TCPAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
