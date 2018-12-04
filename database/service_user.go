package database

import (
	"gopkg.in/mgo.v2"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *Session, dbName string, collectionName string) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection: collection}
}
