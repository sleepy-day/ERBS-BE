package models

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
