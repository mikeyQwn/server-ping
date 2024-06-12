package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`
	PingAddress string `yaml:"ping_address"`
	HidePort    bool   `yaml:"hide_port"`
	TLS         struct {
		Enabled        bool   `yaml:"enabled"`
		CacheDirectory string `yaml:"cache_directory"`
	} `yaml:"tls"`
}

//go:embed config.yaml
var cfgBytes []byte

func ParseConfig() (Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
