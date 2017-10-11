package configuration

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

// Config is the configuration type
type Config struct {
	Dbmigrations dbmigrations
	Logger       zap.Config
	Mysql        MysqlConfig
	SyncEngine   syncEngineConfig
	Server       server
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

type syncEngineConfig struct {
	IntervalInSeconds time.Duration
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
		fmt.Fprintf(os.Stderr, "Error loading config: %s\n", err.Error())
		os.Exit(1)
	}

	return config
}
