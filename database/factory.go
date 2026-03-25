package database

// New - creates a new database instance
func New(configuration *Configuration) *Instance {
	return &Instance{
		configuration: configuration,
	}
}

// ConfigurationNew - creates a new database configuration
func ConfigurationNew() *Configuration {

	// Instance initialization.
	//
	instance := Configuration{}
	instance.Load()

	// Ensure the configuration is valid before
	// returning the instance.
	//
	err := instance.Validate()
	if err != nil {
		panic(err)
	}

	return &instance
}
