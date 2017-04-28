package db

import (
	"github.com/ElectricCookie/das-cms/configLoader"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var connector Connector

// SetConnector sets the db connector to be used
func SetConnector(newConnector Connector) {
	connector = newConnector
}

// Connect connects to the database server
func Connect() {

	var err error

	session, err = mgo.Dial(configLoader.GetConfig().DBHost)

	if err != nil {
		panic(err)
	}

}

// GetDb returns the current db session
func GetDb() Connector {
	return connector
}

// Connector is the interface of all methods a db connector needs to implement
type Connector interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserById(id string) (*User, error)

	GetUserByUsernameOrEmail(input string) (*User, error)

	CreateUser(*User) error

	VerifyEmail(*User) error

	InsertRefreshToken(token *RefreshToken) error

	FindRefreshToken(id string, token string) (*RefreshToken, error)
}
