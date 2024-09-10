package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza App")
	fmt.Println("Please rate our pizza in 1 to 5:")

	//take input from reader
	reader := bufio.NewReader(os.Stdin)
	input, _:= reader.ReadString('\n')

	fmt.Println("Thanks for Rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 yo rating:", numRating+1)
	}
}
