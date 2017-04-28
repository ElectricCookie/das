package mongo

import (
	mgo "gopkg.in/mgo.v2"
)

// Adapter implements the Database interface for MongoDB using m.go
type Adapter struct {
	address string
	session *mgo.Session
}

// Connect to the DB server
func (adapter *Adapter) Connect(address string) {
	adapter.address = address

	session, err := mgo.Dial(address)

	if err != nil {
		panic(err)
	}

	adapter.session = session
}
