package configuration

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/michaelrbond/go-rss-aggregator/errors"
)

// Config is the configuration type
type Config struct {
	Mysql        MysqlConfig
	Server       server
	Dbmigrations dbmigrations
}

type dbmigrations struct {
	Files string
}

// MysqlConfig describes MySQL connection options
type MysqlConfig struct {
	Server   string
	Port     int
	Database string
	User     string
	Password string
}

type server struct {
	Port int
}

// TODO: Is there a better way to do this in Go?
var config Config

// GetConfig returns a configration struct.
func GetConfig() Config {
	if config.Server.Port != 0 {
		return config
	}
	configEnv := os.Getenv("GO_ENV")
	if configEnv == "" {
		configEnv = "local"
	}
	if _, err := toml.DecodeFile(fmt.Sprintf("configuration/%s.toml", configEnv), &config); err != nil {
		errors.Handle(err)
	}
	return config
}

// GetSecret returns the key provided
func GetSecret(s string) string {
	fmt.Printf("Closing Database connection\n")
	return os.Getenv(s)
}
