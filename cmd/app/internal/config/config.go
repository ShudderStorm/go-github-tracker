package config

import (
	"gopkg.in/yaml.v3"
	"io"
)

type Config struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgres"`
	Redis struct {
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
	} `yaml:"redis"`
	GitHub struct {
		Secret string `yaml:"secret"`
	}
}

func Parse(reader io.Reader) (*Config, error) {
	config := &Config{}
	err := yaml.NewDecoder(reader).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
