package main

import (
	"fmt"
	"strings"
)

// concatination in efficient way
func ReverseStr_using_String_Builder(str string)  string {
	var res strings.Builder
	for i := len(str)-1 ; i>=0 ; i-- {
		res.WriteByte(str[i])
	}
	return res.String()
}

func main() {
	fmt.Println("Reverse String")
	s := "Divyansh"

	fmt.Println(ReverseStr_using_String_Builder(s))
	
	i := 0
	j := len(s)-1
	for i <j {
		temp := string(s[i])
		s[i] = string(s[j])
		s[j] = temp
		i++
		j--
	}

	fmt.Println("After reverse : ", s)
}