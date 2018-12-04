package models

//this data struct will hold all the dependencies
//needed to provide our web service
type Data struct {
	PageTitle string
	Data      []User
}

//Users will be our array to keep users in
var Users []User
