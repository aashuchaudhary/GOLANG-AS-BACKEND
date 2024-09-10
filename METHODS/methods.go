package main

import "fmt"

func main() {
	fmt.Println("Welcome to Struct")

	// no Class ,no object, no inheritance , no super ,no parent

	// Utilising the struct
	
	   ashutosh :=User{"Ashu", "ashutosh.chaudhary12@gmail.com", true, 22}
	   fmt.Println(ashutosh)
	// Another way to see it using printf with format specifiers %+v : its just used for entire details
	fmt.Printf("Ashutosh details Are: %+v\n",ashutosh)  
	// for one or two values
	fmt.Printf("Name is %v and email is %v.\n",ashutosh.Name, ashutosh.Email)  
	ashutosh.GetStatus()
	ashutosh.NewMail()
	fmt.Printf("Name is %v and email is %v.\n",ashutosh.Name, ashutosh.Email)  

}


// defining struct

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// plan to import this status

func (u User) GetStatus()  {
fmt.Println("Is user Active: ",u.Status)	
}

// Another function

func (u User) NewMail() {
	u.Email = "xyz@gmailcom"
	fmt.Println("Email of this user: ",u.Email)
}