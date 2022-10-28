package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in golang")

	language := make(map[string] string)
	language["RB"] = "Ruby"
	language["PY"] = "Python"
	language["GO"] = "Golang"

	fmt.Println("map is: ", language)

	delete(language, "GO")
	fmt.Println("map after delete: ", language)

	// loop in golang is interesting
	for key, value := range language{
		fmt.Printf("for key %v value is %v\n", key, value)
	}
}