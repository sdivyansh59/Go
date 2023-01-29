package main

import (
	"fmt"
)

/*
You are given an interger I,
find the absolute value of the interger I.
*/

func main()  {
	fmt.Println("Enter n")
	var n int
	fmt.Scanln(&n)
	
	abs
	fmt.Println(absoluteValue(n))
}


func absoluteValue(n int) int {
	if n <0 {
		return n*-1
	}else {
		return n
	}	
}

