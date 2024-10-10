package app_config

import (
	"log/slog"

	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_FILE_NAME = "config.yml"

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Data struct {
		Dir string `yaml:"dir"`
	}
}

var AppConfig Config

func ReadConfig() {

	// read the config.yaml file
	data, err := os.ReadFile(CONFIG_FILE_NAME)

	slog.Info("read " + CONFIG_FILE_NAME + " file successfully.")

	if err != nil {
		slog.Error(CONFIG_FILE_NAME+" could not be read.", slog.String("error", err.Error()))
	}

	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		slog.Error("unexpected error", slog.String("error", err.Error()))
	}
}
