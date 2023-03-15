package main

import "fmt"

func main() {
	// sum of array using 2 goroutine
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 30}
	channel := make(chan int)
	go sum(arr[:5], channel)
	go sum(arr[5:], channel)
	x, y := <-channel, <-channel
	fmt.Println("sum is: ", x+y)

}

func sum(arr []int, c chan int) {
	var smallSum int
	for _,v := range arr{
		smallSum += v
	} 
	c <- smallSum
}