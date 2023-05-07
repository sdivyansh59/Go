package main

import (
	// "fmt"
	"mylinkedlist/linkedlist"
)

func main() {
	// fmt.Println("Linked list")

	list := linkedlist.LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.PrintList()
	list.Delete(1)
	list.PrintList()

	list2 := linkedlist.LinkedList{}
	list2.Insert("Apple")
	list2.Insert("Banana")
	list2.Insert("Orange")
	list2.PrintList()
	list2.Delete("Orange")
	list2.PrintList()

	// list3 := linkedlist.LinkedList{}
	// list3.Insert(Student{
	// 	name: "Divyansh",
	// 	id: 1,
	// 	active: true,
	// })

	// list3.Insert(Student{
	// 	name: "Ankit",
	// 	id: 2,
	// 	active: true,
	// })

	// list3.Insert(Student{
	// 	name: "Priyanka",
	// 	id: 3,
	// 	active: false,
	// })

	// list3.PrintList()
	
}

// type Student struct {
// 	name string
// 	id int
// 	active bool
// }
// func (s *Student) String () string{
// 	return fmt.Sprintf("%v %v %v ",s.id,s.name, s.active)
// }