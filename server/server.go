// to run this: inside of your go workspace use the command `go run server.go` assuming youre at the proper level in the directory structure
package main

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"
	// Third party packages
	"github.com/Nav31/Food.Now/Models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Get req
	router.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Made a test User
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
	})
	// Post req
	router.POST("/user", func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		user := models.User{}
		json.NewDecoder(req.Body).Decode(&user)
		user.ID = "yello"
		uj, _ := json.Marshal(user)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
	})
	// Delete req
	router.DELETE("/user/:id", func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		w.WriteHeader(200)
	})
	http.ListenAndServe("localhost:3000", router)
}
