package main

import "fmt"
import "github.com/BurntSushi/toml"

// Config is the configuration type
type Config struct {
	Server server
}

type server struct {
	Port int
}

// TODO: Yuck! Should probably be a memoized function?
// Need to research how to do that properly in Go
var config Config

// GetConfig returns a configration struct.
func GetConfig() Config {
	if config.Server.Port != 0 {
		return config
	}
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic("Error opening config file")
	}
	fmt.Printf("%d", config.Server.Port)
	fmt.Printf("%#v\n", config)
	return config
}
