package main

import (
	"fmt"
	"net/http"
	"sync"
)

var singnels = [] string {"test"}

var wg sync.WaitGroup // wg is pointer
var mut sync.Mutex  // pointer


func main() {
	fmt.Println("Goroutines in golang")
	// go greeter("Hello")
	// greeter("world")

	websitelist := [] string {
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(singnels)
}

func getStatusCode (endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOP in endpoint")
	}else{
		mut.Lock()
		singnels = append(singnels, endpoint)
		mut.Unlock()

	}

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}

// func greeter (s string) {
// 	for i := 0; i< 6 ; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(s)
// 	}

// }