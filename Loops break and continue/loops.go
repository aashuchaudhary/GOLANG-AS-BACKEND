package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops ,break and continue")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// For loop

	// for d := 0;d < len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Println("index is %v and value is %d\n",index,day)
	// }

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto abc
		}
		// BREAK AND CONTINUE STATEMENT

		// if rougeValue == 5 {
		// 	break
		// }

		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println("value is: ", rougeValue)
		rougeValue++
	}

	abc :
		fmt.Println("Jumps at Golang.docs")
}
