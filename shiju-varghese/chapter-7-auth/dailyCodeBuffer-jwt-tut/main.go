package main

import (
	"dilycodebuffer/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// fmt.Println("daily code buffer")
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/home", handlers.Home).Methods("GET")
	r.HandleFunc("/refresh", handlers.Refresh).Methods("GET")

	server := &http.Server{
		Addr : ":8080",
		Handler: r,
	}

	fmt.Println("Listening....")
	server.ListenAndServe()

	// fmt.Println()
	// Refresh()
}