package config

import (
	"github.com/spf13/viper"
)

var db Database
var coreDB Database

// EnrolmentDBConfig function for database keys
func EnrolmentDBConfig() Database {
	return db
}

// CoreDBConfig function for database keys
func CoreDBConfig() Database {
	return coreDB
}

// LoadEnrolmentDB function for reading the config
func LoadEnrolmentDB() {
	db = Database{
		Host:     viper.GetString("enrolment.staging.db.host"),
		Username: viper.GetString("enrolment.staging.db.username"),
		Password: viper.GetString("enrolment.staging.db.password"),
		Name:     viper.GetString("enrolment.staging.db.name"),
		Port:     viper.GetInt("enrolment.staging.db.port"),
	}
}

// LoadCoreDB function for reading the config
func LoadCoreDB() {
	coreDB = Database{
		Host:     viper.GetString("core.staging.db.host"),
		Username: viper.GetString("core.staging.db.username"),
		Password: viper.GetString("core.staging.db.password"),
		Name:     viper.GetString("core.staging.db.name"),
		Port:     viper.GetInt("core.staging.db.port"),
	}
}
