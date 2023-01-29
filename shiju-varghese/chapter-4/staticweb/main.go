package main

import (
	"fmt"
	"net/http"
)


func main() {
	fmt.Println("Hello!")
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)


}