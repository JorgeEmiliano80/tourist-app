package config

import (
	"os"
)

type Config struct {
	ServerPort string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUser:     getEnv("DB_USER", "jorgeemiliano"),
		DbPassword: getEnv("DB_PASSWORD", "Jorge41304254#"),
		DbName:     getEnv("DB_NAME", "tourist-app"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
