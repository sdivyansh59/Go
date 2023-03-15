package main

import "fmt"

type Person struct {
	id      int
	name    string
	weight  float32
	address string
}

func (p Person) String() string {
	return fmt.Sprintf("Person Type\n id: %d\n name: %s\n weight: %f\n address: %s\n",p.id, p.name, p.weight, p.address)
}


func main() {

	p1 := Person{
		id: 1, 
		name: "Divyamnsh",
		weight: 50.23, 
		address: "gorakhpur",
	}
	fmt.Println(p1)


}