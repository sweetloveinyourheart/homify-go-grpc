package search_service

import (
	"fmt"
	"homify-go-grpc/internal/search-service/configs"
	"homify-go-grpc/internal/search-service/consumers"
	broker "homify-go-grpc/internal/shared/broker"
	"log"
	"net"
)

func RunSearchServer() {
	configurations := configs.GetConfigs()

	// Kafka consumer setup
	go func() {
		c := consumers.NewSearchConsumer()

		topics := broker.GetTopics()
		c.SubscribeTopics(topics.SearchTopic)
		c.StartSubscribe(topics)

		defer c.CloseConsumer()
	}()

	// Net TCP setup
	_, err := net.Listen("tcp", configurations.TCPAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("🚀 Search Server is listening on port %s ... \n", configurations.TCPAddress)
}
