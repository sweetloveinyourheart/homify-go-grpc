package configs

type Config struct {
	Port                              string
	AuthenticationClientRemoteAddress string
}

func GetConfig() Config {
	return Config{
		Port:                              ":8080",
		AuthenticationClientRemoteAddress: ":50051",
	}
}
