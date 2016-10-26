// to run this: inside of your go workspace use the command `go run server.go` assuming youre at the proper level in the directory structure
package main

import (
	// Standard library packages

	"net/http"

	"gopkg.in/mgo.v2"
	// Third party packages

	"github.com/Nav31/Food.Now/controllers"
	"github.com/julienschmidt/httprouter"
)

func getSession() *mgo.Session {
	// connect to mongoDB
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func main() {
	router := httprouter.New()
	userCtrl := controllers.NewUserController(getSession())
	router.GET("/user/:id", userCtrl.GetUser)
	router.POST("/user", userCtrl.CreateUser)
	router.DELETE("/user/:id", userCtrl.RemoveUser)
	http.ListenAndServe("localhost:3000", router)
}
