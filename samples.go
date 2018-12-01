package main

import "github.com/blavkboy/matcha/models"

//sample user to add to the array
var d = &models.User{
	ID:       "1",
	Username: "The dark one",
	Fname:    "El Pharoah",
	Lname:    "",
	Email:    "akjhdskjahsdkjh@sponges.com",
	Password: "akjsdghhjakhdksjadklashdklad",
	Bio: &models.Bio{
		Caddress: nil,
		Oaddress: &models.Address{
			Street1:  "352",
			Street2:  "Du Toit Street",
			Suburb:   "Wierda Park",
			City:     "Pretoria",
			Province: "Gauteng",
		},
		Sexuality: &models.Sexuality{
			Sex:         models.Male,
			Orientation: models.Hetero,
			Looking:     models.Fun,
			Preferences: nil,
		},
		Hobbies:   nil,
		Interests: nil,
	},
}

//array we will keep users in
var users []models.User
