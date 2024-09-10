package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointer")

	//  var ptr *int
	//  fmt.Println("value of the pointer: ",ptr)

	myNumber := 23
	var ptr = &myNumber // creating a pointer which is reference to something

	fmt.Println("value of actual pointer: ", ptr)  //address
	fmt.Println("value of actual pointer: ", *ptr) // value

	*ptr = *ptr * 2
	fmt.Println("New value is:",myNumber)
}
