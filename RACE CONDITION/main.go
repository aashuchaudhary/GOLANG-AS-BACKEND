//HOW IN THE REAL WORLD SITUATION WE GOINTO RACE CONDITION HOW WE GET INTO AND HOW WE ARE GOING TO SOLVE IT, SO THE RACE CONDITION IS SIMPLY THERE ARE MULTIPLE GO ROUTINES AND WE CAN SAY THERE ARE MULTIPLE THREADS GOING ON AND THIS MEMORY SPACE IS JUST ONE THING IF WE TRY TO WRITE INTO THIS MEMORY SPACE  SUING ONE THREAD AND ANOTHER THREAD  AT THE SAME TIME ITS COME AND SAY I WANT TO WRITE INTO THE SAME MEMORY SPACE OBVIOUSLY WE WILL FACE ISSUES ...

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition detected")

	wg := &sync.WaitGroup{} //pointer
	// mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One  routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two routine")
		mut.Lock() // HERE WE ARE GONNA BE USING THE SAME MEMORY SPACE, SO WE ARE GOING TO GET INTO A RACE CONDITION.
		score = append(score, 2)
		mut.Unlock() // WE ARE GONNA BE USING THE SAME MEMORY SPACE, SO WE ARE GOING TO GET INTO A RACE CONDITION.
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three routine")
		mut.Lock() // HERE WE ARE GONNA BE USING THE SAME MEMORY SPACE, SO WE ARE GOING TO GET INTO A RACE CONDITION.
		score = append(score, 3)
		mut.Unlock() // WE ARE GONNA BE USING THE SAME MEMORY SPACE, SO WE ARE GOING TO GET INTO A RACE CONDITION.
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three routine")
		mut.RLock() //at the time we are writting something in memory the lock is necessary if another threads/routines comes in and write into that we are going into mess  at the time of reading that kind of strict lock is not necessary so anything comes to write in that scurely we lock it but multiple thread comes into the same resourse  and try to read the resource and we are totally allowing that,so here the read write mutex comes in.
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	// ONCE THE WAIT GROUP IS PASSED ON WE HAVE COUPLE OF THINGS TO DO ,I NNEED TO NOTIFY MAIN METHOD THAT I HHAVE ENABLE THE WAIT GROUP.
	wg.Wait()
	fmt.Println(score)
}

//NOTE :EXIT STAUS 66
