package config

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
