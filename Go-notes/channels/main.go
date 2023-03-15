package main

import "fmt"

func routine(s string, c chan string) {
	for i :=0; i<3 ; i++ {
		fmt.Println(i," ",s)
	}
	c <- s+"done"
}

func main() {
	fmt.Println("Go Channels")
	channel := make(chan string)

	go routine("non blocking 1",channel)
	
	go routine("sequencial",channel)

	go routine("non blocking 2",channel)

	message1 := <-channel
	fmt.Println(message1)

	message2 := <-channel
	fmt.Println(message2)

	message3 := <-channel
	fmt.Println(message3)

}