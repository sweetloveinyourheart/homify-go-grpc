package producers

import (
	"fmt"
	broker "homify-go-grpc/internal/shared/broker"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type IPropertyProducer interface {
	InitDeliveryReport()
	ProduceMessages(topic string, value []byte)
	CloseProducer()
}

type PropertyProducer struct {
	producer *kafka.Producer
}

func NewPropertyProducer(configs broker.KafkaConfigs) IPropertyProducer {
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

func (prd *PropertyProducer) ProduceMessages(topic string, value []byte) {
	prd.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
}

func (prd *PropertyProducer) CloseProducer() {
	prd.producer.Close()
}
