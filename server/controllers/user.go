package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nav31/Food.Now/Models"
	"github.com/julienschmidt/httprouter"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an example user
	user := models.User{
		Name:     "Sergy Brin",
		Email:    "sergy@google.com",
		Password: "hello_world",
		Location: "Menlo Park, CA",
		ID:       p.ByName("id"),
	}
	uj, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	user := models.User{}
	json.NewDecoder(req.Body).Decode(&user)
	user.ID = "yello"
	uj, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}
