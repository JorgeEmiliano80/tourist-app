package config

import "os"

var (
	DBHost     = getEnv("DB_HOST", "localhost")
	DBUser     = getEnv("DB_USER", "postgres")
	DBPassword = getEnv("DB_PASSWORD", "Jorge41304254")
	DBName     = getEnv("DB_NAME", "turismo_db")
	DBPort     = getEnv("DB_PORT", "5432")
	ServerPort = getEnv("SERVER_PORT", "8080")
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
