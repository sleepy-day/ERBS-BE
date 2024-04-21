package util

import (
	"fmt"
	"github.com/BurntSushi/toml"
	m "github.com/sleepy-day/ERBS-BE/src/models"
)

func GetDatabaseConfig() m.ErbsConfig {
	var conf m.ErbsConfig

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

func GetConnectionString(conf m.ErbsConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Server,
		conf.Database.Port,
		conf.Database.Name)
}
