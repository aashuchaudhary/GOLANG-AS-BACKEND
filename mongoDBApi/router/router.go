// CREATE ROUTE AND TEST OUR API  WE ARE NOT USING POSTMAN NOR THUNDER CLIENT

package router

import (
	"github.com/aashutoshchaudhary/mongapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router { //return reference

	//ths router is going to be exported and will be using in main file return some type of it
	//create router
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")             //GRAB movie
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")                //create movie
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")          // Update movie
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")         //delete movie
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE") //delete all movies

	return router //this is a mux router that needs to be exported as well as
}
