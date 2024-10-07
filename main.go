package main

import (
	"github.com/BurntSushi/toml"
	"github.com/sleepy-day/ERBS-BE/datafetcher"
)

type ErbsConfig struct {
	Database DatabaseConfig
	Api      ApiConfig
}

type DatabaseConfig struct {
	Server   string
	Port     int
	Name     string
	Username string
	Password string
}

type ApiConfig struct {
	Key string
}

func GetDatabaseConfig() ErbsConfig {
	var conf ErbsConfig

	if _, err := toml.DecodeFile("./erbs-config.toml", &conf); err == nil {
		return conf
	}

	if _, err := toml.DecodeFile("./config/erbs-config.toml", &conf); err == nil {
		return conf
	}

	if _, err := toml.DecodeFile("../config/erbs-config.toml", &conf); err == nil {
		return conf
	}

	panic("Unable to find config file")
}

func main() {
	conf := GetDatabaseConfig()
	datafetcher.FetchData(conf.Api.Key)
}
