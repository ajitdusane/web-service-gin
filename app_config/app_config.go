package app_config

import (
	"log/slog"

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

var AppConfig Config

func ReadConfig() {
	// read the config.yaml file
	data, err := os.ReadFile(configPath)
	slog.Info("read " + configPath + " file successfully.")

	if err != nil {
		slog.Error(configPath+" could not be read.", slog.String("error", err.Error()))
	}

	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		slog.Error("unexpected error", slog.String("error", err.Error()))
	}
}
