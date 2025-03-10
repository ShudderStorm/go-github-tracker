package config

import (
	"gopkg.in/yaml.v3"
	"io"
)

type YamlProvider struct {
	config Config
}

func NewYamlProvider(reader io.Reader) (*YamlProvider, error) {
	config := &Config{}
	err := yaml.NewDecoder(reader).Decode(config)

	if err != nil {
		return nil, err
	}

	return &YamlProvider{config: *config}, nil
}

func (p *YamlProvider) ConfigDatabase() Database {
	return p.config.Database
}

func (p *YamlProvider) ConfigCache() Cache {
	return p.config.Cache
}

func (p *YamlProvider) ConfigOAuth() OAuth {
	return p.config.OAuth
}
