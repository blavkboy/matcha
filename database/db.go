package database

import (
	"gopkg.in/mgo.v2"
)

var dbSession mgo.Session

//StartSession will start our session with mongodb
func StartSession() {
	dbSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	dbSession.SetMode(mgo.Monotonic, true)
}

func GetInstance(collection string) *mgo.Collection {
	c := dbSession.DB("matcha_db").C(collection)
	return c
}
