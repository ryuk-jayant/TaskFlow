package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func InitEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}