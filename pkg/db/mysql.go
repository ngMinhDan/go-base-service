package db

import (
	"database/sql"

	"base/pkg/log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL Configuration Variable
var mysqlCfg DatabaseCfg

// MySQL Variable
var MySQL *sql.DB

// MySQL Connect Function
func mysqlConnect() *sql.DB {
	// Initialize Connection
	conn, err := sql.Open("mysql", mysqlCfg.User+":"+mysqlCfg.Password+"@tcp("+mysqlCfg.Host+":"+mysqlCfg.Port+")/"+mysqlCfg.Name)
	if err != nil {
		log.Println(log.LogLevelFatal, "mysql-connect", err.Error())
	}

	// Test Connection
	err = conn.Ping()
	if err != nil {
		log.Println(log.LogLevelFatal, "mysql-connect", err.Error())
	}

	// Return Connection
	return conn
}
