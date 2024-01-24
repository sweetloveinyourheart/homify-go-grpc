package kafka_configs

type KafkaConfigs struct {
	KafkaServerAddress string
}

type KafkaContexts struct {
	SearchTopic string
	SearchGroup string
}

func GetConfig() *KafkaConfigs {
	return &KafkaConfigs{
		KafkaServerAddress: "localhost:9092",
	}
}

func GetContext() *KafkaContexts {
	return &KafkaContexts{
		SearchTopic: "search",
		SearchGroup: "search-group",
	}
}
