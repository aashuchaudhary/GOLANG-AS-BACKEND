// MAPS OR HASH-TABLE

package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	// Create Maps

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["py"] = "Python"
	languages["c++"] = "C plus plus"
	languages["Rb"] = "Ruby"
	languages["Ms"] = "Mearn Stack"

	fmt.Println("The list of All Languages: ",languages)
	fmt.Println("py shorts for: ",languages["py"])
	
	// Delete Languages
	
	delete(languages,"c++")
	fmt.Println("The list of All Languages: ",languages)

	// Loops in Maps

	for key, value := range languages {
		fmt.Printf("For key %v, Value id %v\n",key,value)
	}
}
