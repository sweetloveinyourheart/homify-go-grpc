package producers

import (
	"fmt"
	kafka_configs "homify-go-grpc/internal/shared/kafka-configs"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type IPropertyProducer interface {
	InitDeliveryReport()
	ProduceMessages(topic string, word string)
	CloseProducer()
}

type PropertyProducer struct {
	producer *kafka.Producer
}

func NewPropertyProducer(configs *kafka_configs.KafkaConfigs) IPropertyProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configs.KafkaServerAddress,
	})

	if err != nil {
		panic(err)
	}

	return &PropertyProducer{
		producer: p,
	}
}

func (prd *PropertyProducer) InitDeliveryReport() {
	for e := range prd.producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}

func (prd *PropertyProducer) ProduceMessages(topic string, word string) {
	prd.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(word),
	}, nil)
}

func (prd *PropertyProducer) CloseProducer() {
	prd.producer.Close()
}
