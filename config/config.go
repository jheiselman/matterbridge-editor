package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// Config represents a matterbridge-web configuration
type Config struct {
	Web          *web
	Matterbridge *matterbridge
}

type web struct {
	IPAddress   string `toml:"address"`
	Port        int    `toml:"port"`
	ContextPath string `toml:"context"`
}

type matterbridge struct {
	ConfigPath string `toml:"config"`
}

// ReadConfig reads a configuration file provided by path and returns a config object
func ReadConfig(path string, logger *logrus.Logger) *Config {
	config := &Config{}
	if _, err := toml.DecodeFile(path, config); err != nil {
		logger.Warning("Error decoding configuration file:", err)
		return nil
	}

	return config
}
