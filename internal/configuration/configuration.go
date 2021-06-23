package configuration

import (
	"gopkg.in/yaml.v2"
	"os"

	"github.com/pkg/errors"
)

func New(path string) (*Config, error) {
	configuration := &Config{}
	if path == "" {
		return nil, errors.New("path empty")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&configuration)
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}

	return configuration, err
}
