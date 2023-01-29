package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("add without goroutine")
    start := time.Now()

	sum := add (10000)
	elapsed := time.Since(start)
	fmt.Printf("sum is: %v and time: %v\n",sum, elapsed)

	// cuncurrent add
	 start2 := time.Now()
	sum2 := addConcurrent(10000)
	elapsed2 := time.Since(start2)
	fmt.Printf("sum is: %v and time: %v",sum2, elapsed2)
	

}

func add (n int) int {
	sum := 0
	for i:= 0 ; i<= n ; i++ {
		sum += i
	}
	return sum
}


func addConcurrent(n int) int {
	// utilise all cores on machine
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)

	var sum int64
	max := n
	sizeOfParts := max / numOfCores

	var wg sync.WaitGroup

	for i := 0; i< numOfCores ; i++ {
		// Divide the input into parts
		start := i * sizeOfParts
		end := start + sizeOfParts

		// run each parth as seperate goroutine

		wg.Add(1)
		go func (s int, e int, *sum int ) {
			defer wg.Done()

			var partSum int
			
			// calculate sum for each part
			for i := s; i< e; i++{
				partSum += i 
			}
			&sum = &sum+partSum 
		}(start, end, &sum)

	}

	fmt.Println(sum)
}