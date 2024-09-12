package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: Error loading .env file")
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		Port:       os.Getenv("PORT"),
	}

	config.DatabaseURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	return config
}
