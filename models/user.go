package models

import (
	"context"
	"time"

	"github.com/blavkboy/matcha/database"
	"github.com/blavkboy/matcha/mlogger"
	"gopkg.in/mgo.v2/bson"
)

//The user file here will deal with all structures and
//methods that will work on the user and have to deal with
//the user's part in the database

//User struct will represent a user and will server
//as a storage mechanism for the server to keep track
//of user data.
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Fname    string        `json:"fname" bson:"fname"`
	Lname    string        `json:"lname" bson:"lname"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

func NewUser(user *User) {
	mlogger := mlogger.GetInstance()
	err, client := database.InitDB()
	if err != nil {
		mlogger.Println("Error: ", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		mlogger.Println("Error: ", err)
	}
	collection := client.Database("matcha").Collection("users")
	res, err := collection.InsertOne(context.Background(), bson.M{"username"})
}
