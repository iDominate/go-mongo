package main

import (
	"log"
	"net/http"

	"github.com/iDominate/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	userController := controllers.NewUserController(getSession())
	r.GET("/user/:id", userController.GetUser)
	r.POST("/user/", userController.CreateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	http.ListenAndServe("localhost:8000", r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("MongoDB Session Creation Failed. :", err)
	}

	return session
}
