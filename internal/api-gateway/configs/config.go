package configs

type Config struct {
	Port string
}

func GetConfig() Config {
	return Config{
		Port: ":8080",
	}
}
