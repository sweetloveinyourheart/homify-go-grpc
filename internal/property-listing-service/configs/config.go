package configs

type PropertyListingConfig struct {
	TCPAddress string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func GetConfig() PropertyListingConfig {
	return PropertyListingConfig{
		TCPAddress:       ":50052",
		PostgresHost:     "localhost",
		PostgresPort:     "50042",
		PostgresUser:     "root",
		PostgresPassword: "root",
		PostgresDatabase: "homify-property-listing",
	}
}
