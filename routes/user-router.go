package routes

import (
	"crud/controllers"
	"github.com/gorilla/mux"
)


/*
	register router endpoints package
 */
var RegisterUserRoutes = func(router *mux.Router){

	router.HandleFunc("/user/",controllers.Create).Methods("POST")
	router.HandleFunc("/user/{id}",controllers.Update).Methods("PUT")
	router.HandleFunc("/user/{id}",controllers.Delete).Methods("DELETE")
	router.HandleFunc("/user/{id}",controllers.Get).Methods("GET")
}
