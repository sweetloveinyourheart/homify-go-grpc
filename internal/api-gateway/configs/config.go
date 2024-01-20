package configs

type Config struct {
	Port                               string
	AuthenticationClientRemoteAddress  string
	PropertyListingClientRemoteAddress string
}

func GetConfig() Config {
	return Config{
		Port:                               ":8080",
		AuthenticationClientRemoteAddress:  ":50051",
		PropertyListingClientRemoteAddress: ":50052",
	}
}
