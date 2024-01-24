package configs

type SearchServiceConfig struct {
	TCPAddress string
}

func GetConfig() *SearchServiceConfig {
	return &SearchServiceConfig{
		TCPAddress: ":50053",
	}
}
