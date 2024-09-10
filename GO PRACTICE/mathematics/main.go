package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Mathematics....")
// 	var myNumberOne int = 2
// 	var myNumberTwo float64 = 4.5

// 	fmt.Println("The sum is: ", myNumberOne + int(myNumberTwo)) // not valid syntax

// for maths 
func Intn(n int) int
// Intn returns,as an int ,a non-negative pseudo-random number panics if n<=0.

//crypto
func Int
func Int(rand io.Reader,max *big.Int)(n *big.Int,err error)
//int returns a uniform random value in [0, max).It panics if max<= 0.

rand.Seed(time.Now().UnixNano())
fmt.Println(rand.Intn(5)+1)

myRandomNumber, _ := rand.Int(rand.Reader,big.NewInt(5))
fmt.Println("The Random number is: ", myRandomNumber)
}
