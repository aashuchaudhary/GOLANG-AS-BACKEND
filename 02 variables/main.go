package main

import "fmt"

// jwttoken := 30000 // its not allowed outside the method its global. the correct way to declare cariable as same as the convertion nal method.

// var jwttoken int = 30000 // allowed

// DECLARING CONST

const LoginToken string = "ghbhjvjg" //here logintoken start with  capital letter means it has some significance that it is Public varaible now. in golang we simply write it as start wit capital letter as it consider as public

func main() {
	// fmt.Println("Variables")

	// STRINGS

	// var username string = "Ashutosh"
	// fmt.Println(username)
	// fmt.Printf("variable is of type: %T \n", username)

	// BOOLEAN

	// var isLoggedIn bool = false
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// SMALL FLOAT

	// var smallFloat float64 = 255.8757646854
	// fmt.Println(smallFloat)
	// fmt.Printf("Variable is of type: %T \n", smallFloat)

	// SMALL INT
	 
	// var smallVal uint8 = 255
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is of type: %T \n", smallVal)

	// DEFAULT VALUES AND SOME ALLIASE

	// var anotherVariable int  // jumst initialise it but dont put any value get result zero as expected.
	// fmt.Println(anotherVariable)
	// fmt.Printf("Variable is of type: %T \n",anotherVariable)

	// IMPLICIT TYPE: ANOTHER WAY TO DECLARE VARIABLE

	// var website = "LearnCoading.in"
	// fmt.Println(website)

	// No VAR STYLE TYPE : its allowed inside only method

	numberOfUser := 30000.0
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n",LoginToken)
}
