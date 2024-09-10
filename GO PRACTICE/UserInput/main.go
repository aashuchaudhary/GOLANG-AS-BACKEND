package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Welcome to user input"
	fmt.Println(Welcome)
	//wes eee eroor here becoz we dont use it Anywhere else
	reader := bufio.NewReader(os.Stdin) // Read from stdin
	fmt.Println("Enter the rating of our Pizza:")

	//whwenever raeder reads we stored into  a variable and the introduce comma okay and errror okay.
// fpr error we can simply use underscore here
	input, _ := reader.ReadString('\n')
	fmt.Println("Thaanks for Rating", input)

	fmt.Printf("Type odf this rating :%T", input)
	//its possible when we reading something from stdin there might be a chance oto get something wrong the there might be chance to get something error that self.

	
}
