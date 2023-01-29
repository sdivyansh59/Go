package main

import "fmt"

type  Person struct {
	Id int
	Name string
	Age int
}

func (p Person) string() string {
	return fmt.Sprintf("Person Type\n Id : %d\n Name : %s \n Age : %d \n",p.Id, p.Name, p.Age)
}

func main() {
	fmt.Println("Hello")

	var l = Person{1, "Divyansh", 20}
	fmt.Println(l)

}