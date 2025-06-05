package config

type Config struct {
	DBConfig  DBConfig  `toml:"db"`
	JWTConfig JWTConfig `toml:"jwt"`
}

type DBConfig struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	Port     string `toml:"port"`
	SSLMode  string `toml:"ssl_mode"`
	Timezone string `toml:"timezone"`
}

type JWTConfig struct {
	SecretKey string `toml:"secret_key"`
}
