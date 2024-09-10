// POST REQUEST

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to post request server")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

// Get request

func PerformGetRequest() {
	// content url
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

func PerformPostJsonRequest() {
	// content url
	const myurl = "http://localhost:9000/post"

	// create some of the data : fake json payload
	// content body

	requestBody := strings.NewReader(`
	{
	"coursename":"Lest's go with golang",
	"price":"0",
	"platform":"LearnCodeOnline.in"
	}
	`)

	// post request

	response, err := http.Post(myurl, "application/json", requestBody)
	// error handling
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

// perform post request

func PerformPostFormRequest() {
	// content url
	const myurl = "http://localhost:9000/postform"

	//create  form data
	data := url.Values{} // data comming from url and it can be access by url package and we have. values inside and initially create an emplty value pair.

	data.Add("firstname", "ashutosh")
	data.Add("lastname", "chaudhary")
	data.Add("email", "ashutosh.chaudhary12@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
