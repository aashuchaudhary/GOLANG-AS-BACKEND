// go mod is a tool

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// Running in the server
	// http.ListenAndServe(":4000",r) //this code so some errors in the case of web we can doclassic comma syntaxbut we are presenting another package which is log whichave a option of fatal.
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod user")
}

func serveHome(w http.ResponseWriter, r *http.Request) { //w is http response writer and r is request type pointer
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
