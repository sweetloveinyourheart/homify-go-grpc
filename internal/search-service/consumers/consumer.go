package consumers

import (
	"fmt"
	broker "homify-go-grpc/internal/shared/broker"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ISearchConsumer interface {
	SubscribeTopics(topics ...string)
	StartSubscribe(topics broker.KafkaTopics)
	CloseConsumer()
}

type SearchConsumer struct {
	client *kafka.Consumer
}

func NewSearchConsumer() ISearchConsumer {
	configs := broker.GetConfigs()
	groups := broker.GetGroups()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KafkaServerAddress,
		"group.id":          groups.SearchGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return &SearchConsumer{
		client: c,
	}
}

func (csm *SearchConsumer) SubscribeTopics(topics ...string) {
	csm.client.SubscribeTopics(topics, nil)
}

func (csm *SearchConsumer) StartSubscribe(topics broker.KafkaTopics) {
	for {
		msg, err := csm.client.ReadMessage(time.Second)
		if err == nil {
			topic := msg.TopicPartition.Topic

			switch topic {
			case &topics.SearchTopic:
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			default:
				fmt.Println("Unknown message")
			}

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
