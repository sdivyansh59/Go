package main

import "fmt"

type Student struct {
	id int
	name string
}

func main() {

	m := make(map[Student]bool)

	obj1 := Student{
		id: 1,
		name: "Divyansh",
	}
	obj2 := Student{
		id: 2,
		name: "Ankit",
	}

	obj3 := Student{
		id: 1,
		name: "Divyansh",
	}

	m[obj1] = true
	m[obj2] = true
	m[obj3] = false
	// fmt.Printf("%v %p",obj1, &obj1)
	// fmt.Printf("%v %p",obj3, &obj3)
	fmt.Println(m)

}