package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

}

// Note -> defer statement will be push into stack and it will be executed 
// at lst ( think like this )

func myDefer(){
	for i := 0; i <5 ;i++ {
		defer fmt.Println(i)
	}
}

