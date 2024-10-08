package app_config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "config.yml"

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

// create a config struct and deserialize the data into that struct
var AppConfig Config

func ReadConfig() {
	// read the config.yaml file
	data, err := os.ReadFile(configPath)

	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		panic(err)
	}
}
