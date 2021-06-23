package configuration

import "time"

type Config struct {
	Server `yaml:"server"`
}

func (c *Config) GetHTTP() Server {
	return c.Server
}

type Server struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout struct {
		Server time.Duration `yaml:"server"`
		Write  time.Duration `yaml:"write"`
		Read   time.Duration `yaml:"read"`
		Idle   time.Duration `yaml:"idle"`
	} `yaml:"timeout"`
}
