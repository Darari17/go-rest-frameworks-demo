package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var AppConfig Config

func init() {
	data, err := os.ReadFile(".toml")
	if err != nil {
		log.Fatalf("failed to read config.toml: %v", err)
	}

	if err := toml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalf("failed to unmarshal toml: %v", err)
	}
}
