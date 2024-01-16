package configs

type AuthenticationConfig struct {
	TCPAddress string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	JwtSecret string
}

func GetConfig() AuthenticationConfig {
	return AuthenticationConfig{
		TCPAddress:       ":50051",
		PostgresHost:     "localhost",
		PostgresPort:     "50041",
		PostgresUser:     "root",
		PostgresPassword: "root",
		PostgresDatabase: "homify-authentication",
		JwtSecret:        "secret",
	}
}
