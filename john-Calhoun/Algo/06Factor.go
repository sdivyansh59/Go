package main

import "fmt"

func factorOfNum(primes []int, num int) []int {
	if (num ==0 || n==1 ){
		return []
	}
	return []

}

var primeList []int

func main() {
	primeList = append(primeList,2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37, 39, 41, 43)
	fmt.Println(primeList)
	factors := factorOfNum(primeList, 30)
	fmt.Println(factors)
}