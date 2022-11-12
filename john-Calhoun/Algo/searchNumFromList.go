package main

import "fmt"


func sumRecursion( input [] int) int{
	if len(input) ==0 {
		return 0
	}
	return input[0] + sumRecursion(input[1:])
}

// check stack over flow
func recurFun (n int) int {
	return n + recurFun(n-1)
}

func main() {
	fmt.Println("Search num from list")
	slice := [] int {1,20,3,45,500,-0,75,111}
	// slice := [] int {}
    target := 500
	sum := 0

	for _, val := range slice {
		// fmt.Println(idx, val)
		if val == target {
			fmt.Println("Found")
		}
		sum += val
	}

	// sum of list
	fmt.Println("Sum of list is: ", sum)

	// fmt.Println(sumRecursion(slice))
	
   fmt.Println(recurFun(1))
	
}