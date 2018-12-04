package models

//The user file here will deal with all structures and
//methods that will work on the user and have to deal with
//the user's part in the database

//User struct will represent a user and will server
//as a storage mechanism for the server to keep track
//of user data.
type User struct {
	ID       string
	Username string
	Fname    string
	Lname    string
	Email    string
	Password string
}
