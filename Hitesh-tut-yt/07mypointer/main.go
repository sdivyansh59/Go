package main

import "fmt"

func main() {
	fmt.Println("Hello pointer")

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr2 = & myNumber
	
	fmt.Println("Value of actual pointer is ", ptr2)
	fmt.Println("Value of actual pointer is ", *ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("New value is : ", myNumber)


}