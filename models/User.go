package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	UserName string             `bson:"userName,omitempty"`
	Email    string             `bson:"email,omitempty"`
}
