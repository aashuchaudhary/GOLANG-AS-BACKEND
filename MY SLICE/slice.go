package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices: ")

	var fruitList = []string{"Apple","Banana","Mango"}
	fmt.Printf("Types of fruitLists is %T\n", fruitList)

	// Adding some new data

	fruitList = append(fruitList, "Peas","Lichi")
	fmt.Println("The new Updated list will be: ",fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Make Keyword is used yo create slices and maps

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 535
	highScores[2] = 676
	highScores[3] = 977


	highScores = append(highScores,857,564,897)

	fmt.Println(highScores)

	// Sorting we use Sort keyword

	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))  // Boolean Values and give true it its is sorted

	// Remove values from Slices Based Index

	var course = []string{"React js","Python","Ruby","Next js","Django","Flutter","Rust"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index],course[index + 1:]...)
	fmt.Println(course)

}
