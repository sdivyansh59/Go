package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	result := add(1,3)
	fmt.Println(result)


	// we can pass any no. of arguments in proAdder 
	//becz it is using slice internally

	proResult , mymessage:= proAdder(1,2,5)
	fmt.Println(proResult, mymessage)


}

func proAdder(values... int) (int, string) {
	total := 0

	for val := range values {
		total += val
	}

	return total , "Hi I'm pro adder function "
}

func add( first int , second int) int{
	return first +second
}