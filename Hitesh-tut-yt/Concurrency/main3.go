package main

import "fmt"

func even(ch chan int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			ch <- i
		}
	}
}

func odd(ch chan int) {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			ch <- i
		}
	}
}

func main() {
	fmt.Println("even odd")
	chEven := make(chan int)
	// chOdd := make(chan int)

	go even(chEven)
	go odd(chEven)

	for  {
		fmt.Println(<-chEven)
		// fmt.Println(<-Even)
	}

}