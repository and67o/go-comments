package configuration

type Config struct {
	HTTPServer
}

func (c *Config) GetHTTP() HTTPServer {
	return c.HTTPServer
}

type HTTPServer struct {
	Host string `mapstructure:"server_host"`
	Port string `mapstructure:"server_port"`
}