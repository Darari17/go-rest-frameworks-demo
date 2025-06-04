package helpers

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYAMLConfig[T any](path string, target *T) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, target); err != nil {
		return err
	}

	return nil
}
