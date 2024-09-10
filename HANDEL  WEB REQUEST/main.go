package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("Lco web request")
	// http request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // Caller's responnsibility to close the connection.

	// Reading Respose
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// convert datatype into content

	content := string(databytes)
	fmt.Println(content)
}
