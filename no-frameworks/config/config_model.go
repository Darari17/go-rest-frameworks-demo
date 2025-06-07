package config

type Config struct {
	AppConfig AppConfig
	DBConfig  DBConfig
	JWTConfig JWTConfig
}

type AppConfig struct {
	AppName string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBTimezone string
}

type JWTConfig struct {
	JWTSecretKey string
}
