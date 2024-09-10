package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// world, one, two
// 0, 1, 2, 3, 4
// hello, 43210, two, one, world


// Defer means :The Function Execution Happens line by line, and the defer will be just above the closing curley brases.
// the o/p for result will bw like  last in first out.


func myDefer()  {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	
}