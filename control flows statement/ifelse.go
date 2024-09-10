package main

import "fmt"

func main() {
	fmt.Println("Welcome to control Flow Statement")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular user"

	} else if loginCount > 10 {
		result = "Watch Out"
	}else {
		result = "Exactly 10 login Count"
	}
	fmt.Println(result)

	// EVEN ODD

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	}else {
		fmt.Println("Number is odd")
	}

	// INITIALISING AND ASSIGINIG VALUE AT SAME TIME

	if num := 3;num <10 {
		fmt.Println("Num is less than 10")
	}else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	// }

}
