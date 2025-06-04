package helpers

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYAMLConfig(path string, target interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, target)
}
