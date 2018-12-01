package core

import "github.com/mongodb/mongo-go-driver/mongo"

//GetDatabase will return an instance of the Database
func GetDatabase() {
	client, err := mongo.NewClient()
}
