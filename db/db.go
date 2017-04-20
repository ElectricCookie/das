package db

import (
	"github.com/ElectricCookie/das-cms/configLoader"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// Connect connects to the database server
func Connect() {

	var err error

	session, err = mgo.Dial(configLoader.GetConfig().DBHost)

	if err != nil {
		panic(err)
	}

}

// Disconnect from the current session
func Disconnect() {
	if session != nil {
		session.Close()
	}
}

// GetDb returns the current db session
func GetDb() *mgo.Database {
	return session.DB("das")
}
