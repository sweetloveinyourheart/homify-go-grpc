package configs

type Config struct {
	Port                        string
	AuthClientRemoteAddress     string
	PropertyClientRemoteAddress string
}

func GetConfig() Config {
	return Config{
		Port:                        ":8080",
		AuthClientRemoteAddress:     ":50051",
		PropertyClientRemoteAddress: ":50052",
	}
}
