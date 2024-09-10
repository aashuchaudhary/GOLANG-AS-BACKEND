package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Array")

	// how to initialise the Array
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Cabbage", "Peas"}
	fmt.Println("vegetable list is : ", vegList)
	fmt.Println("vegetable list is : ", len(vegList))

}
