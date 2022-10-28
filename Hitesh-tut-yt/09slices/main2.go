package main

import "fmt"

func main() {
	// how to remove value from slice

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var idx int = 2
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println(courses)
}
