package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Specification represents structured configuration variables
type Specification struct {
	Port      string `envconfig:"SERVICE_PORT" default:"50051"`
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
