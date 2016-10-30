package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Test model
type Test struct {
	ID  bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Foo string        `json:"foo"`
}
