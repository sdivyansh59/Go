package main

import "fmt"


func main() {
	channel := make(chan int, 10)

	go fibonacci(cap(channel), channel)

	for i := range channel {
		fmt.Println(i)
	}
}


func fibonacci(n int, ch chan int ){
	x,y  := 0,1
	for i:= 0 ; i<n ; i++ {
		ch <- x
		x,y = y, x+y
	}
	close(ch) // imp to close
}

