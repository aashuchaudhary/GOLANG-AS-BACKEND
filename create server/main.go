// make get request ,post request  but the data in the json format and post form request but the data willbe encoded url forms.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:9000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("The status code : ", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength)

	//  createing a builder
	var responseStrings strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseStrings.Write(content)

	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseStrings.String())

	// fmt.Println(content) //content is in byte format
	// fmt.Println(string(content))

	// Another method to create the get request using String package as it has  builder method , as bulder is used to efficiently build a string using write method
}
