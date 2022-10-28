package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("conversion among data types")
	fmt.Println("please rate out pizza b/t 1-5")

	reader := bufio.NewReader(os.Stdin);

	input , _ := reader.ReadString('\n');

	fmt.Println("Thanks for rating ", input)

	numberRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("numberRating  = ", numberRating+1)
		//, numberRating+1
	}




}