package broker

type KafkaConfigs struct {
	KafkaServerAddress string
}

func GetConfigs() KafkaConfigs {
	return KafkaConfigs{
		KafkaServerAddress: "localhost:9092",
	}
}
