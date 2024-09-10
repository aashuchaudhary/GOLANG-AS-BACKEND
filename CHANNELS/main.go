// CHANNELS ARE THE WAY THAT MULYIPLE GO ROUTINE TALK TO EACH OTHER.

// CHANNELS ARE USED TO SEND AND RECEIVE VALUES BETWEEN GOROUTINES.

// CHANNELS CAN BE CREATED USING THE make FUNCTION.

// CHANNELS ARE BLOCKING, MEANS THE GOROUTINE WILL NOT EXECUTE UNTIL THE OTHER GOROUTINE SENDS A VALUE ON THE CHANNEL.

// CHANNELS CAN BE CLOSED WITH THE CLOSE FUNCTION.

// CHANNELS CAN BE USED TO SEND AND RECEIVE VALUES BETWEEN GOROUTINES.

// SELECT STATEMENT IS USED TO SELECT THE GOROUTINE THAT WILL BE READY TO EXECUTE. SELECT STATEMENT CAN BE USED WITH CHANNELS AND TIMEOUTS.

// SELECT STATEMENT IS USED TO SELECT THE GOROUTINE THAT WILL BE READY TO EXECUTE. SELECT STATEMENT CAN BE USED WITH CHANNELS AND TIMEOUT
package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Welcome to channnels in Golang")

	//creating a channel
	// CHANNELS ARE THE WAY OR A KIND OF PIPELINE THROUGH MULTIPLE GOROUTINE INTERACT EACH OTHER AND WE NEED TO DEFINE WHAT KIND OF VALUES SHOULD BE PASSING WITH EACH OTHER IT WOULD BE INT OR STRING

	// myCh := make(chan int)
	myCh := make(chan int, 2)  //buffer channnel: it tell adding one more value int it but it dont going to consume it but it dont show any error.
	wg := &sync.WaitGroup{}

	/* myCh <- 5 //push values in channel and we use arrow here and the arrow is pointing towards left in golang
	fmt.Println(<-myCh)    // Receive value from channel */

	wg.Add(2)

// recieve only 
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// to find the 0 is of closed channel or message fromsingnals 

		val,  isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<- myCh)  //LISTENER ,each lisetener take individual values if there is two value and one listener then we get error to handle it we need to either two listener or one listener ,when we have onelistener  then  that listener get into a loop and trying to get all the values in that cas eweuse buffer channel 
		
		// fmt.Println(<- myCh)

		
		wg.Done()
	}(myCh, wg)

	// send only

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)

		// myCh <- 5
		// myCh <- 6

	/* 	myCh <- 0 //we can get 0 by two reason either by passinng zero as asignal value  or a msg from channel that you are lsitining to a closed channel.
 */
		/* close(myCh)  //what if we pass the value after closing the channel and and now we try to read and then it go into a panic mode and then we have to design a methods which accept the situation which we are reciving a special error we are sending toa close channel. */

		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
