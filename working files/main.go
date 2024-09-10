package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	// content to put in files and handdle it 
	content  := "This need sto go in a filr - golang.dev"

	// put some text file and work with that  by os module or os package

	file, err := os.Create("./mylcofile.txt")  // file is created

	// checking error, error have something 

	// if err != nil {
	// 	panic(err)   //panic keyword throws an eror will shutdown execution of program and show you waht error you are faceing
	// }
		checkNillErr(err)


	// for  writing
	len, err := io.WriteString(file, content)
		checkNillErr(err)

	fmt.Println("lenght is: ",len)
	defer file.Close()

	readFile("./mylcofile.txt")
}

// For Reading  the File 

func readFile(filename string)  {
	// ioutils is used to read the file 

	// whenever we read the file it's not  been read in the string format, and the data we read form the intenet is in the byte format.

	databyte,err:= ioutil.ReadFile(filename)
	// checking error
	checkNillErr(err)
	fmt.Println("Text data inside the file is \n",databyte)
	fmt.Println("Text data inside the file is \n",string(databyte))
	// For handling data 
	
}

// here wer used error checking many times so in place of that we create a function for error checking

func checkNillErr(err error)  {
	if err != nil {
		panic(err)  
	}
}
