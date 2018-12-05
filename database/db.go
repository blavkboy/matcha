package database

import (
	"fmt"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"gopkg.in/mgo.v2"
)

var dbSession mgo.Session

//StartSession will start our session with mongodb
func SaveUser(collection string, obj *models.User) {
	logger := mlogger.GetInstance()
	dbSession, err := mgo.Dial("localhost")
	defer dbSession.Close()
	if err != nil {
		panic(err)
	}
	dbSession.SetMode(mgo.Monotonic, true)
	c := dbSession.DB("matcha_db").C(collection)
	fmt.Println("Got collection: ", c)
	err = c.Insert(obj)
	if err != nil {
		logger.Println("Error: ", err)
	}
}
