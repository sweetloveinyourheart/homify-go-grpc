package configs

type SearchServiceConfigs struct {
	TCPAddress string
}

func GetConfigs() SearchServiceConfigs {
	return SearchServiceConfigs{
		TCPAddress: ":50053",
	}
}
