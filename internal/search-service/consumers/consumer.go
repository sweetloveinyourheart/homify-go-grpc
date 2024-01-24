package consumers

import (
	"fmt"
	kafka_configs "homify-go-grpc/internal/shared/kafka-configs"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ISearchConsumer interface {
	Subscribe(topics ...string)
	CloseConsumer()
}

type SearchConsumer struct {
	client *kafka.Consumer
}

func NewSearchConsumer(configs *kafka_configs.KafkaConfigs, contexts *kafka_configs.KafkaContexts) ISearchConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KafkaServerAddress,
		"group.id":          contexts.SearchGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return &SearchConsumer{
		client: c,
	}
}

func (csm *SearchConsumer) Subscribe(topics ...string) {
	csm.client.SubscribeTopics(topics, nil)

	for {
		msg, err := csm.client.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func (csm *SearchConsumer) CloseConsumer() {
	csm.client.Close()
}
