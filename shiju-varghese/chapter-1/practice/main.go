package main

import (
	"calculatorApp/exception"
)
func main() {
	// fmt.Println("calculator app ")
	// fmt.Println("Enter 1st and 2nd val :")
	// var  x,y int
	// fmt.Scan(&x, &y )
	// fmt.Printf("The x = %d and y = %d \n", x, y)

	// fmt.Println(calc.Add(x,y))
	// fmt.Println(calc.Substract(x,y))

	// swapchar.SwapCharactersInSentence("This is Divyansh! ")

	// arr := [5]int{}
	// fmt.Printf("%T",arr)
	// fmt.Println(len(arr), cap(arr))

	// for _,v := range arr {
	// 	fmt.Println(v)
	// }

	// fmt.Println("Starting to panic.")
	// exception.DoPanic()
	// fmt.Println("Program regains control after recover")

	// ans1, err1 := exception.Divide(10,5)

	// ans2, err2 := exception.Divide(10,0)

	// fmt.Println(ans1,err1,ans2, err2)

	exception.CheckDeferSe3quence()
}

