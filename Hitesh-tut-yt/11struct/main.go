package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang ; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 37}

	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are %+v \n", hitesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}