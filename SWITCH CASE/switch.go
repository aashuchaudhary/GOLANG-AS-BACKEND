package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case ")

	// Dice Game, rand Using math
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) +  1 
	fmt.Println("Value of dice is",diceNumber)

	switch diceNumber {
	case 1: 
	fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice move to spot 2")
	case 3:
		fmt.Println("Dice move to spot 3")
		fallthrough
	case 4:
		fmt.Println("Dice move to spot 4")
		fallthrough  // used to moke in next case
	case 5:
		fmt.Println("Dice move to spot 5")
	case 6:
		fmt.Println("Dice move to spot 6 and roll again")
	default:
		fmt.Println("What was that!")
  }
}
