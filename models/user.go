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
	Location GeoLocation   `json:"location" bson:"Location"`
	Profile  Profile       `json:"profile" bson:"profile"`
}

//constant for the cost
var Cost int = 6

//NewUser iss effectively how a new user is registered onto the
//system with their username, email and password. All other details
//are only necessary when the user model design is decided and users
//can provide more information about themselves.
func NewUser(user *User) *User {
	mlogger := mlogger.GetInstance()
	//Either initialize the database or get an instance of it
	client := database.GetInstance()
	defer client.Close()
	c := client.DB("matcha").C("users")
	index := mgo.Index{
		Key:        []string{"username", "email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		mlogger.Println("Error: ", err)
		panic(err)
	}
	mlogger.Println("Ensured Index")
	user.ID = bson.NewObjectId()
	err = c.Insert(&user)
	if err != nil {
		mlogger.Println("Error: ", err)
		return nil
	}
	mlogger.Println("Inserting User")
	user = FindUser("username", user.Username)
	return user
}

func UpdateUser(updatedUser *User) *User {
	logger := mlogger.GetInstance()
	var orig User
	user := FindUser("id", updatedUser.ID)
	client := database.GetInstance()
	c := client.DB("matcha").C("users")
	err := c.Find(bson.M{
		"_id": user.ID,
	}).One(orig)
	if err != nil {
		logger.Println("Could not update user: ", err)
		return nil
	}
	logger.Println("Updatating user: ", orig)
	defer client.Close()
	return user
}

//FindUser will return a User struct of the user being queried
//based on key value pair of the caller's choosing. Still needs
//to be tested extensively.
func FindUser(key string, value interface{}) *User {
	body := new(User)
	mlogger := mlogger.GetInstance()
	client := database.GetInstance()
	defer client.Close()
	c := client.DB("matcha").C("users")
	err := c.Find(bson.M{
		key: value,
	}).One(body)
	if err != nil {
		mlogger.Println("Error: ", err)
		return nil
	}
	return (body)
}
