package models

type ErbsConfig struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Server   string
	Port     int
	Name     string
	Username string
	Password string
}
