package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6)+1
	fmt.Println("Val of dice is: ", diceNum)

	switch diceNum {
	case 1: 
		fmt.Println("Dice value is 1 ")
	case 2:
		fmt.Println("Dice value is 2 ")
	case 3:
		fmt.Println("Dice value is 3 ")
	case 4:
		fmt.Println("Dice value is 4 ")
	case 5:
		fmt.Println("Dice value is 5 ")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6 ")
	default:
		fmt.Println("What was that ! ")					

	}

	//Note -> break is not required 
	// Note -> fallthrough means all case after that will also be executed
  // What is seed ?
}