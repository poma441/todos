package config

type Config struct {
	ServerHost string
	ServerPort string
	DbHost     string
	DbPort     string
	DbUsername string
	DbName     string
	DbPassword string
	DbSslMode  string
}

func InitConfig() *Config {
	config := new(Config)

	// Инициализация конфига

	return config
}
