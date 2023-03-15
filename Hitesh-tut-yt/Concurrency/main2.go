package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
    start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch<- result{url, err,0}
	}else {
		t := time.Since(start).Round(time.Microsecond)
		ch<- result{url,nil,t}
		resp.Body.Close()
	}
}


func main() {

	list := [] string{
		"https://youtube.com",
		"https://nytimes.com",
		"https://linkedin.com",
		"https://google.com",
	}

	results := make(chan result)

	for _,v := range list {
		go get(v,results)
	}

	// loop result ch only 
	for range list {
		r := <-results

		if r.err != nil {
			fmt.Printf("%-20s %s\n", r.url,r.err)
		}else {
			fmt.Printf("%-20s %s\n", r.url,r.latency)
		}
	}
		
}