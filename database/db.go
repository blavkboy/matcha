package database

import (
	"sync"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var once sync.Once
var connection *mongo.Client

func InitDB() (error, *mongo.Client) {
	mlogger := mlogger.GetInstance()
	var err error = nil
	once.Do(func() {
		client, err := mongo.NewClient("mongodb://localhost:27017")
		connection = client
		if err != nil {
			mlogger.Println("Error: ", err)
		}
	})
	if err != nil {
		return err, nil
	}
	return nil, connection
}
