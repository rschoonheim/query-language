package database

// Configuration - holds the configuration for the database
type Configuration struct {
	Name string `yaml:"name"`
}

// Load - loads the configuration for the database
func (c *Configuration) Load() {

}

// Validate - validates the configuration for the database
func (c *Configuration) Validate() error {
	return nil
}
