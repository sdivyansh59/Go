package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	size int
}

func main() {
	fmt.Println("LLinked list")

	mylist := linkedList{}

	// Doubt what node1 will store if I will not  mention &
	// ex: node1 := node{data: 48}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}

	mylist.addLast(node1)
	mylist.addFront(node2)
	mylist.addLast(node3)

	fmt.Println(mylist)



}



func (l *linkedList) addLast (n *node) {

	temp := l.head
	if temp == nil {
		temp = n
		l.size++
		return
	}

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = n
	l.size++

}


func (l *linkedList) addFront (n *node) {
	temp := l.head
	l.head = n
	l.head.next = temp

	l.size++
}