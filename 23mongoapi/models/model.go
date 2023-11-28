package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string `json:"moviename,omitempty"`
	Watched bool `json:"isWatched,omitempty"`
}