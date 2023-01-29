package main

import "fmt"

func main() {
	// buffered channel
	ch := make(chan int,6)
	go func() {
		for i := 0; i < 6; i++ {
			// send iterator over channel
			fmt.Println("sending ", i)
			ch <- i
		}
		close(ch)
	}()

	// range over channel to receive value
	for v := range ch {
		fmt.Println("receiving: ",v)
	}
}