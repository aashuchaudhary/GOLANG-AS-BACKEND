package main

import "fmt"

func main() {
	fmt.Println("Welcome to Struct")

	// no Class ,no object, no inheritance , no super ,no parent

	// Utilising the struct
	
	   ashutosh :=User{"Ashu", "ashutosh.chaudhary12@gmail.com", true, 22}
	   fmt.Println(ashutosh)
	// Another way to see it using printf with format specifiers %+v : its just used for entire details
	fmt.Printf("Ashutos details Are: %+v\n",ashutosh)  
	// for one or two values
	fmt.Printf("Name is %vand email is %v.",ashutosh.Name, ashutosh.Email)  

}


// defining struct

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
