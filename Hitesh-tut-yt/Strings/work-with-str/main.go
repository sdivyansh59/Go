package main

import (
	"fmt"
	"strings"
)

/*
Pattern Counting Problem: Count the number of times that a pattern
appears in a longer text.
	Input: Strings text and pattern.
	Output: The num of times that patterns as a substing of text

*/

func main () {
	fmt.Println(strings.Count("banana", "an"))
	fmt.Println(strings.Count("CGATATAGGATAC", "ATA"))
	// strings.Count() doesn't account for overlapping occurences of pattern

	fmt.Println("Hello "+ "World!")

	s := "Hi there!"
	fmt.Println("The length if s is: ", len(s))
	fmt.Println("The first symbol of s is: ", s[0], string(s[0]))
	fmt.Println("The last symbol of s is: ", s[8],string(s[8]))

	if s[1] == 'i' {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	
 
	// substring
	fmt.Println(s[3:5])
	fmt.Println(s[:3])
}