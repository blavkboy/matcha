package models

import (
	"github.com/blavkboy/matcha/database"
	"github.com/blavkboy/matcha/mlogger"
	mgo "gopkg.in/mgo.v2"
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
	//Either initialize the database or get an instance of it
	err, client := database.InitDB()
	if err != nil {
		mlogger.Println("Error: ", err)
		return
	}
	c := client.DB("matcha").C("users")
	index := mgo.Index{
		Key:        []string{"username", "email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		mlogger.Println("Error: ", err)
		panic(err)
		return
	}
	mlogger.Println("Ensured Index")
	err = c.Insert(&user)
	if err != nil {
		mlogger.Println("Error: ", err)
		panic(err)
	}
	mlogger.Println("Inserting User")
	/*
		res, err := collection.InsertOne(ctx, bson.M{
			"username": user.Username,
			"fname":    user.Fname,
			"lname":    user.Lname,
			"email":    user.Email,
			"password": user.Password,
		})
		if err != nil {
			mlogger.Println("Error: ", err)
			return
		}
		id := res.InsertedID
		mlogger.Println("Insertion ID: ", id)
		mlogger.Println("User: ", user)
	*/
}
