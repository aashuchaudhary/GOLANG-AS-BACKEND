package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aashutoshchaudhary/mongapi/router"
)

func main() {
	fmt.Println("Mongodb Api")
	// bring router here , but this time its not comming from mux router this time its comming from router
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listing at port 4000 ...")

}
