package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "WELCOME TO USER INPUT"
	fmt.Println(welcome)

	// we see error here  becz we are not use it anywhere.

	reader := bufio.NewReader(os.Stdin) //stdin means standard input output

	fmt.Println("Enter the rating for our pizza:")

	// whatever the reader reads we stored into a variable and  there introduce comma ok and error ok.
	// for error we can use simply underscore.

	// comma ok // err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
	fmt.Printf("Type of this ratting is :%T", input)

	// its possible when  we reading something from standard input the might we chance to get something wrong the there might be chance to get somthing error that  self.

	// for that we stored error as
	// input, err := reader.ReadString('\n')  // standard syntax of golang

	// comma ok and // err ok is like try and catch.
}
