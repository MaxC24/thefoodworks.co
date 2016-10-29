package models

import "gopkg.in/mgo.v2/bson"

type (
	// User is the user for app
	User struct {
		Id       bson.ObjectId `json:"id" bson:"_id"`
		Name     string        `json:"name" bson:"name"`
		Email    string        `json:"email" bson:"gender"`
		Password string        `json:"password" bson:"age"`
	}
)
