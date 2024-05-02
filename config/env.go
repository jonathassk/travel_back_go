package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type DbConfig struct {
	DbName string
	DbUser string
	DbHost string
	DbPort int
	Region string
}

var Envs = initConfig()

func initConfig() DbConfig {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found")
		return DbConfig{}
	}
	return DbConfig{
		DbName: getEnv("DB_NAME", "postgres"),
		DbUser: getEnv("DB_USER", "postgres"),
		DbHost: getEnv("DB_HOST", "localhost"),
		DbPort: getEnvInt("DB_PORT", 5432),
		Region: getEnv("REGION", "sa-east-1"),
	}
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		stringResult, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return stringResult
	}
	return fallback
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
