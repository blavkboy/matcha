package database

import (
	"github.com/blavkboy/matcha/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userModel struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password string
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newUserModel(u *models.User) *userModel {
	return &userModel{
		Username: u.Username,
		Password: u.Password,
	}
}

func (c *userModel) toModelUser() *models.User {
	return &models.User{
		ID:       c.Id.Hex(),
		Username: c.Username,
		Password: c.Password,
	}
}
