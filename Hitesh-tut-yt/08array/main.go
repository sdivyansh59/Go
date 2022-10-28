package main

import "fmt"

func main() {

	var fruitsList [4]string

	fruitsList[0] = "apple"
	fruitsList[1] = "Banana"
	fruitsList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitsList)
	fmt.Println("Fruit list is : ", len(fruitsList))

	var vegList = [5] string {"Potato", "Tomato", "Beans"}
	fmt.Println("vegList is: ", vegList)
	fmt.Println("Size of vegList is: ", len(vegList))

}