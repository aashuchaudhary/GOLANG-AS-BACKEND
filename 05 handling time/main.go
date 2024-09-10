package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	
	fmt.Println(presentTime)

	// Standard time

	fmt.Println(presentTime.Format("01-02-2006 15:04:04 Monday")) //syntax or  format for order of date time and day by golang it is standard.

	// Reverse syntax in which we dont want to current time we want to create the time

	createDate := time.Date(2020, time.December, 10, 23, 23, 0, 0, time.UTC)  // for location we choose time.UTC
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-06 Monday"))
}


// creatiing environment for different system :first check go env in powershell.

// $env:GOOS="linux"   :Sets the target operating system to Linux.

// $env:GOARCH="amd64"   :Sets the target architecture to amd64 (64-bit architecture).

// $env:CGO_ENABLED="0"  :Disables cgo (Cgo allows Go packages to call C code), which simplifies cross-compilation.

// go build -x : Compiles the Go application and provides verbose output (-x flag) to show the commands being executed during the build process. This can help in debugging if something goes wrong.

