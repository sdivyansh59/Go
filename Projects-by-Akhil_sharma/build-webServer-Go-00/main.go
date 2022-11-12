package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not suppported", http.StatusNotFound)
		return
	}

	fmt.Println(w,"hello!")
}


func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm() err : %v", err)
		return
	}
	fmt.Println(w,"Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n", name)
	fmt.Fprintf(w,"Address = %v", address)

}

func main() {
	// fmt.Println("My First web server in GO lang\n")

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
		
	
}