package main

import "fmt"

func main() {
	fmt.Println("if else for golang")

	logicCount := 23
	var result string

	if logicCount < 10 {
		result = "Regulae user"
	}else if logicCount > 10 {
		result = "watch out"
	}else {
		result = "exactely 10"
	}

	fmt.Println(result)

	// we can define var in if stataement as well
	// it is useful when we receive sone valu from api 
	if num := 3 ; num <10 {
		fmt.Println("Num is less than 10")
	}else {
		fmt.Println("Num is greate r than 10")
	}

	// we use nil 
	
	// if err != nil{
	// 	fmt.Println(err)
	// }

}