package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 and 5:")

	// take input from reader

	reader := bufio.NewReader(os.Stdin) // its not going to do anything its just a reference
	input, _ := reader.ReadString('\n') //for properly read we can hold it into a reference  and all the result in into a variable.

	fmt.Println("Thnaks for ratting: ", input)
	numRating, err  := strconv.ParseFloat(strings.TrimSpace(input),64)  //converting string into float

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Added 1 to you Rating: ",numRating+1)
	}
}
