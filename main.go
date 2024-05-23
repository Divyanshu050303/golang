package main

import (
	"fmt"
	"net/http"

	"golang-syntax/curd_operation/controller"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	// syntax.DatType();
	// syntax.Variables();
	//  syntax.Loops()
	// syntax.StatementSs();
	// syntax.ArrayOperation();

	// syntax.Input()

	// syntax.Gomap();

	// syntax.GoSlice()
	fmt.Println("Hello world!")

	r := httprouter.New()
	uc := controller.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		panic(err)
	}
	return s
}
