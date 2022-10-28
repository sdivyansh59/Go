package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitsList = []string{"apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruits lis is: %T\n", fruitsList)
	fruitsList = append(fruitsList, "Mangos", "Banana")

	fmt.Println("fruits list aftr append is : ", fruitsList)


	fruitsList = append(fruitsList[1:])
	fmt.Println(fruitsList)

	// usindg make()to create array
	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 20
	highScores[2] = 30
	highScores[3] = 40

	fmt.Println(highScores)

	// it will not give error ( unlike array using new ())
	highScores = append(highScores, 50,60)
	fmt.Println(highScores)

	// sort
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}