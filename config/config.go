package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string
	DatabaseURI   string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &Config{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
		DatabaseURI:   viper.GetString("DATABASE_URI"),
	}
}
