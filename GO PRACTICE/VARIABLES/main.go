/* // variables
package main

import (
	"fmt"
)

// jwttoken := 3000; // not allowed
// var jwttoken int = 3000 // allowed
// Declaring const
const LoginToken string = "sfdvnjnfd" // for public vatiable we star with capital letter

func main() {
	fmt.Println("hello")
	//Strings
	var Username string = "Ashutosh"
	fmt.Println(Username)
	fmt.Printf("Variable is of type: %T \n", Username)

	//Boolean

	var IsLoggedIn bool = true
	fmt.Println(IsLoggedIn)
	fmt.Printf("Variavle uos of type: %T \n", IsLoggedIn)

	//Small float

	var smallFloat float64 = 223.5547457568
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T\n", smallFloat)

	//small int
	var smallval uint8 = 255
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type %T  \n", smallval)

	// Default values and some alliance
	var anptheVariable int
	fmt.Println(anptheVariable)
	fmt.Printf("variable is of type %T \n ", anptheVariable)

	//implicit variable

	var website = "LearnCoding.com"
	fmt.Println(website)

	//no var style type : its allowed only methods

	numberOfUsers := 345356
	fmt.Println(numberOfUsers)
	fmt.Printf("variable is pf type %T\n", numberOfUsers)
	fmt.Println(LoginToken)
	fmt.Printf("variavble is of type %T \n ", LoginToken)
}
 */
 

package main

import (
    "fmt"
)

var insufficientfundmessage string= "not money"
var sufficientfundmessage string= "have money"
var accountbalance float64= 100.0
var bulkmeassage float64= 75.0
var isPremiumUser bool= false
var discountamount float64=0.10
var finalcost float64

finalcost = bulkmeassage
if isPremiumUser{
	finalcost = bulkmessage  * (1 - discountamount) 
}

if accountbalance >=finalcost{
	accountbalance-=finalcost
	fmt.Println("purchase sucess")
}
else{
	fmt.Println("purchase dont sucess")
}

fmt.Println("Account Balance",accountbalance)