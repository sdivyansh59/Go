package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// video: https://www.youtube.com/watch?v=RfF4lk9loXI
func randomNum () int{
	return rand.Intn(100)
}
func displayMesaage(ch chan string) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i< 10 ; i++ {
		go func ()  {
			defer wg.Done()
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("This is value %d in channel ",randomNum())
		}()
		
	}
	wg.Wait()
	close(ch)
}
func main() {
	ch := make(chan string,3)
	// special variable that is used to communicate b/w goroutine
	// ch <- "This is value 1 in channel"
	// ch <- "This is value 2 in channel"
	// fmt.Println(<-ch)
	// close(ch)

	go displayMesaage(ch)
	for chv := range ch {
		fmt.Println(chv)
	}
	
	
	


}