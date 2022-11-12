package main

import "fmt"

func fibonacii(n int) int {
	if n == 0 || n == 1 {
		// fmt.Println(n)
		return n
	}
	res := (fibonacii(n-1) + fibonacii(n-2))
	fmt.Println(res)
	return res
}


func main() {
	n := 5
	fmt.Println("fibonacii of num")
	fibonacii(n)
}