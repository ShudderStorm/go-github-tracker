package config

import (
	"bufio"
	"gopkg.in/yaml.v3"
)

type YamlProvider struct {
	reader *bufio.Reader
}

func NewYamlProvider(reader bufio.Reader) *YamlProvider {
	return &YamlProvider{reader: &reader}
}

func (p *YamlProvider) Parse(filename string) (*Config, error) {
	config := &Config{}
	err := yaml.NewDecoder(p.reader).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
