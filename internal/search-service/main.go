package search_service

import (
	"fmt"
	"homify-go-grpc/internal/search-service/configs"
	"homify-go-grpc/internal/search-service/consumers"
	kafka_configs "homify-go-grpc/internal/shared/kafka-configs"
	"log"
	"net"
)

func RunSearchServer() {
	configurations := configs.GetConfig()
	kafkaConfigs := kafka_configs.GetConfig()
	kafkaContexts := kafka_configs.GetContext()

	// Kafka consumer setup
	go func() {
		c := consumers.NewSearchConsumer(kafkaConfigs, kafkaContexts)
		c.Subscribe(kafkaContexts.SearchTopic)
		defer c.CloseConsumer()
	}()

	// Net TCP setup
	_, err := net.Listen("tcp", configurations.TCPAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("🚀 Search Server is listening on port %s ... \n", configurations.TCPAddress)
}
