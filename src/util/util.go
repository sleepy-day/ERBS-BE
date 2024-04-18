package util

import (
	"fmt"
	"github.com/BurntSushi/toml"
	m "github.com/sleepy-day/ERBS-BE/src/models"
)

func GetDatabaseConfig() m.ErbsConfig {
	var conf m.ErbsConfig
	if _, err := toml.DecodeFile("./config/config.toml", &conf); err != nil {
		panic("Error opening config file")
	}

	return conf
}

func GetConnectionString(conf m.ErbsConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Server,
		conf.Database.Port,
		conf.Database.Name)
}
