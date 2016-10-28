package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nav31/Food.Now/Models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser gets a particular user
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(404)
	// 	return
	// }
	oid := bson.ObjectIdHex(id)
	user := models.User{}
	fmt.Println(oid)
	if err := uc.session.DB("Food_Now").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser makes a user
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user := models.User{}
	json.NewDecoder(req.Body).Decode(&user)
	user.Id = bson.NewObjectId()
	uc.session.DB("Food_Now").C("users").Insert(user)
	uj, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}
