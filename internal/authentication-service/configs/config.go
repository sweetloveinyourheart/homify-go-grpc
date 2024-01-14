package configs

type AuthenticationConfig struct {
	TCPAddress string
}

func GetConfig() AuthenticationConfig {
	return AuthenticationConfig{
		TCPAddress: ":50051",
	}
}
