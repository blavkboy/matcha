package database

import (
	"os"
	"sync"

	"github.com/blavkboy/matcha/mlogger"
	mgo "gopkg.in/mgo.v2"
)

//Decided to rewrite my whole approach to using the database

var once sync.Once
var session *mgo.Session

func InitDB() (error, *mgo.Session) {
	mlogger := mlogger.GetInstance()

	var err error = nil
	once.Do(func() {
		session, err = mgo.Dial("mongodb://localhost:27017")
		if err != nil {
			mlogger.Println("Error: ", err)
			os.Exit(1)
		}
		session.SetMode(mgo.Monotonic, true)
	})
	if err != nil {
		return err, nil
	}
	mlogger.Println(session)
	return nil, session
}
