package db

import (
	"base/pkg/log"

	mgo "gopkg.in/mgo.v2"
)

// Mongo Configuration Variable
var mongoCfg DatabaseCfg

// MongoSession Variable
var MongoSession *mgo.Session

// Mongo Variable
var Mongo *mgo.Database

// Mongo Connect Function
func mongoConnect() (*mgo.Session, *mgo.Database) {
	// Initialize Connection
	conn, err := mgo.Dial(mongoCfg.User + ":" + mongoCfg.Password + "@" + mongoCfg.Host + ":" + mongoCfg.Port + "/" + mongoCfg.Name)
	if err != nil {
		log.Println(log.LogLevelFatal, "mongo-connect", err.Error())
	}

	// Test Connection
	err = conn.Ping()
	if err != nil {
		log.Println(log.LogLevelFatal, "mongo-connect", err.Error())
	}

	// Set Connection to Monotonic
	conn.SetMode(mgo.Monotonic, true)

	// Return Connection
	return conn, conn.DB(mongoCfg.Name)
}
