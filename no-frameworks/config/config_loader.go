package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Gagal load file .env: %v", err)
	}

	return &Config{
		AppConfig: AppConfig{
			AppName: os.Getenv("APP_NAME"),
			AppPort: os.Getenv("APP_PORT"),
		},
		DBConfig: DBConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBPort:     os.Getenv("DB_PORT"),
			DBSSLMode:  os.Getenv("DB_SSLMODE"),
			DBTimezone: os.Getenv("DB_TIMEZONE"),
		},
		JWTConfig: JWTConfig{
			JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
	}
}
