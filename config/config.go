package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config is the structure representing the database information needed to connect.
type Config struct {
	Server   string
	Database string
}

// Read decodes the config.toml file to create a Config data value.
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
