package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions.")
	greeter() // Passing Reference

	// Note :We cant define a function inside a function.

	// func greeterTwo()  {
	// 	fmt.Println("Anoter method")
	// }
	// greeterTwo()

	//  Creating a Special method  that add two  numbers
	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	// proResult,_ := proAdder(2,5,8,7,6,9,3,1)
	// fmt.Println("Pro result is: ",proResult)
	// or
	proResult,mymessage:= proAdder(2,5,8,7,6,9,3,1)
	fmt.Println("Pro result is: ",proResult)
	fmt.Println("Pro message is: ",mymessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// Dont know how much value will come but add all of them

// func proAdder(values ...int) int {
// 	total := 0

// 	for _, val := range values {
// 		total += val
// 	}
// 	return total

// }

func proAdder(values ...int) (int,string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hey Pro result functions"
}

func greeter() {
	fmt.Println("Namastey from Golang Group")
}
