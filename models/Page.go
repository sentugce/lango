package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Page struct {
	ID       primitive.ObjectID `bson:"email,omitempty"`
	Name     string             `bson:"Name,omitempty"`
	UserName string             `bson:"UserName,omitempty"`
	Desc     string             `bson:"Desc,omitempty"`
	URI      string             `bson:"URI,omitempty"`
}
