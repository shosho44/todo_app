package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server server
}

type server struct {
	Port string `yaml:"port"`
}

func Load() (*Config, error) {
	yamlByte, err := os.ReadFile("internal/configs/config.yaml")

	if err != nil {
		return nil, fmt.Errorf("calling ReadFile: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlByte, &config.Server)

	if err != nil {
		return nil, fmt.Errorf("calling Unmarshal: %w", err)
	}
	return &config, nil
}
