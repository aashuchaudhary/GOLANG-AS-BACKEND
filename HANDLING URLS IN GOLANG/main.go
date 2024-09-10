package main

import (
	"fmt"
	"net/url"
)

// creating urls
const myurl string = "httpd://lco.dev:3000/learn?coursename=reactjs&paymentid=ghgbj456" //here & is like comma in urls
// const myurl string  = "https://google.com"

func main() {
	fmt.Println("Welcome to handel urls in golang")
	fmt.Println(myurl)

	// Parsing urls : means extracting some information
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// Storing  result.RawQuery isnide a variables.

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// constructing urls from information

	partsofurls := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=ashutosh",
	}

	anotherURL := partsofurls.String()
	fmt.Println(anotherURL)
}
