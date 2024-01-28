package broker

type KafkaTopics struct {
	SearchTopic string
}

type KafkaGroups struct {
	SearchGroup string
}

func GetTopics() KafkaTopics {
	return KafkaTopics{
		SearchTopic: "search",
	}
}

func GetGroups() KafkaGroups {
	return KafkaGroups{
		SearchGroup: "search-group",
	}
}
