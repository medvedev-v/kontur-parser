package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DefaultInputFilename  string   `yaml:"defaultinputfilename"`
	DefaultOutputFilename string   `yaml:"defaultoutputfilename"`
}

var appConfig *Config

// LoadConfig загружает конфигурацию из YAML файла
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	appConfig = &config
	return appConfig, nil
}

// GetConfig возвращает загруженную конфигурацию
func GetConfig() *Config {
	return appConfig
}
