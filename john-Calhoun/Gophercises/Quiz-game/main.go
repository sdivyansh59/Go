package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Quiz Game")
	// open file 
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv val using csv.Reader
	csvReader := csv.NewReader(f)
	// read line by line
	correct, wrong := 0 ,0
    quesNo := 1

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		correctResponse, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Panic("not able to convert inot int")
		}
		fmt.Printf("Question %d is: %v \n", quesNo, rec[0])
		fmt.Println("Enter your response: ")
		var ans int
        fmt.Scanf("%d\n",&ans)
		
		// check response
		if ans == correctResponse {
			correct ++
		}else {
			wrong++
		}

		quesNo++
	}

	fmt.Printf("Correct %d wrong %d ", correct, wrong)
}