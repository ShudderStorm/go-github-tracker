package config

type Config struct {
	Database Database `yaml:"database"`
	Cache    Cache    `yaml:"cache"`
	OAuth    OAuth    `yaml:"oauth"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Cache struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type OAuth struct {
	ClientID string `yaml:"client_id"`
	Secret   string `yaml:"secret"`
}
