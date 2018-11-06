package config

import (
	"github.com/spf13/viper"
)

// Database struct for connection
type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

var db Database

// DB function for database keys
func DB() Database {
	return db
}

// LoadDB read the config
func LoadDB() {
	db = Database{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Name:     viper.GetString("db.name"),
	}
}
