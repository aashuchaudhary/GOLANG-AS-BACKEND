// IN PREVIOSUS WE TRAET JSONAS A STRING AND NOW WE SEE HOW TO CREATE AND HANDEL JSON AS WEBREQUEST.
// encoding of json means :having data can be array,slices  and convet data into Valid json

package main

import (
	"encoding/json"
	"fmt"
)

// creating structure
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` //slice of string
}

func main() {
	fmt.Println("Welcome to learn something more about json files.")
	// EncodingJson()
	DecodeJsonData()
}

func EncodingJson() {
	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "Multiple", "abs123", []string{"web-dev", "js"}},
		{"Nodejs Bootcamp", 199, "Multiple", "ahfj123", []string{"backend-dev", "js"}},
		{"Nextjs Bootcamp", 399, "Multiple", "ajgfh123", nil},
	}

	// Package this data as json data
	// marshal is like implementation amd we have to pass a interface (it is alternative version of struct)

	// finalJson,err := json.Marshal(lcoCourses)

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // marshal indent takes two parameter first is the interface and the second one is the basedon which what we indent the values.
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

// Decode  Json Data

func DecodeJsonData() {
	jsonDataFromWeb := []byte(`
	  {
                "coursename": "Reactjs Bootcamp","Price": 299,"website": "Multiple","tags": ["web-dev","js"]
      }
	 `)

	//  what ever the data comes from the web and we create a structure and put them into the structure
	var LcoCourse course
	// WANT TO CHECK THE FAKE JSON DATA IS VALID DOR NOT
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was Valid")
		json.Unmarshal(jsonDataFromWeb, &LcoCourse) // passing INTERFACE AND RFERENCE
		fmt.Printf("%#v\n", LcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//  some cases where you just want to adddata to key value pair
	// REMEMBER WHEN WE CREATE A MAP FOR ONLINE JSON STRING THE FIRST VALUE IS KEY VALUE PAIR IS GAURENTEED BUT THE VALUE IS NOT GUARENTED IT CAN BE INTEGER,STRING OR AN ARRAY.
	var onlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlinedata)
	fmt.Printf("%#v\n", onlinedata)

	for k, v := range onlinedata {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
