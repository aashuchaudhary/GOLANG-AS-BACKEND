package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Weelcome to time study")

	PresentTime := time.Now()
	fmt.Println(PresentTime)

	//standard time

	fmt.Println(PresentTime.Format("201-02-2006 15:04:04 Monday"))

	CreateDate := time.Date(2020, time.December, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(CreateDate)
	fmt.Println(CreateDate.Format("01-02-06 Monday"))
}
