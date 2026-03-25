package database

import (
	"sync"

	"gopkg.in/yaml.v3"
)

// New - creates a new database instance
func New(configuration *Configuration) *Instance {
	return &Instance{
		configuration: configuration,
		state:         StateStarting,
		stateMutex:    sync.RWMutex{},
	}
}

// ConfigurationFromYaml - creates a new database configuration from yaml bytes
func ConfigurationFromYaml(configurationBytes []byte) (*Configuration, error) {

	var configuration Configuration

	err := yaml.Unmarshal(configurationBytes, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
