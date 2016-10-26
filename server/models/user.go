package models

import "gopkg.in/mgo.v2/bson"

type (
	// This is will be our user
	User struct {
		Name     string        `json:"name" bson: "name"`
		Email    string        `json:"email" bson: "email"`
		Password string        `json:"password" bson: "password"`
		Location string        `json:"location" bson: "location"`
		Id       bson.ObjectId `json:"id" bson: "_id"`
	}
)
