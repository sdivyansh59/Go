package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Learn String in Go")
	var c byte = 'H'
	fmt.Println("c is: ",c)  // is 'H' printed ?
	mj := string(45)
	fmt.Println("mj is: ",mj) // is "45" printed?



	/*
	Go represent symbol as Number
	and string as arrays of symbols
	When Go sees the string "Hello World!" it visualizes it 
	as an arrray, [72 101 108 108 111....]
	the correspondence between symbol and number is called ASCII
	*/

	/*
	Recall: os.Args stores an array of all command Line parameters
	as an array of string
	*/


	// x := os.Args[1]
	// y := os.Args[2]
    // // to run type : go run main.go Hello World!
	// fmt.Println("x+y is: ", x+y)


	// strconv allows us to convert between types involing string nature
	mj = strconv.Itoa(45)
	fmt.Println(mj)

	kj := "10"
	fmt.Println(kj)
	kjj, err1 := strconv.Atoi(kj)
	if err1 != nil {
		log.Panic(err1)
	}
	fmt.Println(kjj)
}