package main

import (
	"log"
	"net/http"
)

const message = "Hello GopherCon UK 2018!"

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalf("Server fail to start")
	}

}