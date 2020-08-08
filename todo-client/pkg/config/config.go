package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Specification represents structured configuration variables
type Specification struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS" default:"localhost:50051"`
	ClientPort    string `envconfig:"CLIENT_PORT" default:"3000"`
}

// LoadEnv loads config variables into Specification
func LoadEnv() (*Specification, error) {
	var conf Specification
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
