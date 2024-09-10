// Textio users are reporting that we're billing them for wildly inaccurate amounts. They're supposed to be billed .02 dollars for each text message sent.

//SOLUTION :

// package main

// import "fmt"

// func main() {
// 	messagesFromDoris := []string{
// 		"You doing anything later??",
// 		"Did you get my last message?",
// 		"Don't leave me hanging...",
// 		"Please respond I'm lonely!",
// 	}
// 	numMessages := float64(len(messagesFromDoris))
// 	costPerMessage := .02

// 	// don't touch above this line

// 	totalCost := costPerMessage * numMessages

// 	// don't touch below this line

// 	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
// }