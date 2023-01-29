package main

import (
	"calculator/calc"
	"fmt"
)

func main() {
	fmt.Println("This is Calculator!")
	var x,y int = 10,5
	fmt.Println(calc.Add(x,y))
	fmt.Println(calc.Substract(10,5))
}

/*
Notes: package name should be ralative to module name
*/
