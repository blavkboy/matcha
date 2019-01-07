package models

import "gopkg.in/mgo.v2/bson"

type Profile struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	UserID     bson.ObjectId `json:"userid" bson;"userid"`
	ProfilePic []byte        `json:"profilepic" bson:"profilepic"`
}
