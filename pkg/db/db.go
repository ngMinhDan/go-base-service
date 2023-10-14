/*
Package db: Provides methods for working with Database.
Package Functionality: Connects to Database: PostgreSQL, Mysql, Mongo

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package db

import (
	"base/pkg/config"
	"strings"
)

// Define a DATABASE (Postgres, MySQL, MongoDB...) Configuration Struct
type DatabaseCfg struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// Define a method "checkConfig" for the "DatabaseCfg" struct
// Check config's info is complete or not
func (db DatabaseCfg) checkConfig() bool {

	if len(db.Host) != 0 && len(db.Port) != 0 &&
		len(db.User) != 0 && len(db.Password) != 0 &&
		len(db.Name) != 0 {
		return true
	} else {
		return false
	}
}

// Define a method "get" for the "DatabaseCfg" struct
// Check config's info is complete or not
func (db *DatabaseCfg) getConfig() {
	db.Host = config.Config.GetString("DB_HOST")
	db.Port = config.Config.GetString("DB_PORT")
	db.User = config.Config.GetString("DB_USER")
	db.Password = config.Config.GetString("DB_PASSWORD")
	db.Name = config.Config.GetString("DB_NAME")

}

// Initialize Function in DB
// Get type of database driver (postgresql, nysql, mongodb...)
// Connect by info written in config's file
func init() {
	// Get type of database driver
	switch strings.ToLower(config.Config.GetString("DB_DRIVER")) {
	case "postgres":
		// Set default port for PostgreSQL
		config.Config.SetDefault("DB_PORT", "5432")

		// get config from file
		psqlCfg.getConfig()

		// check config's info
		if psqlCfg.checkConfig() == true {
			// Do PostgreSQL Database Connection
			PSQL = psqlConnect()
		}
	case "mysql":
		// Set default port for Mysql
		config.Config.SetDefault("DB_PORT", "3306")

		// get config from file
		mysqlCfg.getConfig()

		// check config's info
		if mysqlCfg.checkConfig() == true {
			// Do MySQL Database Connection
			MySQL = mysqlConnect()
		}

	case "mongo":
		// Set default port for MongoDB
		config.Config.SetDefault("DB_PORT", "27017")

		// get config from file
		mongoCfg.getConfig()

		// check config's info
		if mongoCfg.checkConfig() == true {
			// Do Mongo Database Connection
			MongoSession, Mongo = mongoConnect()
		}
	}
}
