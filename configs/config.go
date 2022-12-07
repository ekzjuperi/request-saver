package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct contains service config.
type Config struct {
	Port   string `yaml:"port"`
	DBPath string `yaml:"dbPath"`
}

// GetConfig get config from .yaml file.
func GetConfig() (*Config, error) {
	f, err := os.Open("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
