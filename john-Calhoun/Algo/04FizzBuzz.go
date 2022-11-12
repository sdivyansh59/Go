package main

import (
	"fmt"
	"strconv"
)

func printFizzBuzz(n int) string {
	res := ""
	if n%3 == 0 {
		res += "Fizz"
	}
	if n%5 == 0 {
		res += "Buzz"
	} else {
		
		res += strconv.Itoa(n)
	}
	return res
}

func main() {
	n := 16
	for i := 0; i < n; i++ {
		fmt.Println(printFizzBuzz(i))

	}

}