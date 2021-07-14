package configuration

import "time"

type Config struct {
	Server `yaml:"server"`
	DBConf `yaml:"db_conf"`
	Auth   `yaml:"auth"`
	Redis  `yaml:"redis"`
}

func (c *Config) GetHTTP() Server {
	return c.Server
}

func (c *Config) GetDBConf() DBConf {
	return c.DBConf
}

func (c *Config) GetAuth() Auth {
	return c.Auth
}

func (c *Config) GetRedis() Redis {
	return c.Redis
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

type Auth struct {
	RefreshKey    string        `yaml:"refresh_key"`
	AccessKey     string        `yaml:"access_key"`
	AccessExpire  time.Duration `yaml:"access_expire"`
	RefreshExpire time.Duration `yaml:"refresh_expire"`
}
type Redis struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Password string `yaml:"password"`
}

type DBConf struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
}