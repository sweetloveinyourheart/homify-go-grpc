package broker

type KafkaTopics struct {
	SyncProperties string
}

type KafkaGroups struct {
	SearchGroup string
}

func GetTopics() KafkaTopics {
	return KafkaTopics{
		SyncProperties: "sync-search",
	}
}

func GetGroups() KafkaGroups {
	return KafkaGroups{
		SearchGroup: "search-group",
	}
}
