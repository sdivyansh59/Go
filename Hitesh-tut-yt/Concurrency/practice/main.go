package main

import (
	"fmt"
	"sync"
	"time"
)



func main() {
	var data int

	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		data++
		
	}()
	wg.Wait()
	fmt.Println(data)
}

func fun(str string) {
	// 
	for i := 0; i < 3; i++ {
		fmt.Println(str)
		time.Sleep(1 * time.Millisecond)
	}
}