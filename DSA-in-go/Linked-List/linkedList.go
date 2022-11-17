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
	myLinkedList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	node4 := &node{data: 40}
	node5 := &node{data: 50}

	myLinkedList.addFront(node1)
	myLinkedList.addFront(node2)
	myLinkedList.addFront(node3)
	myLinkedList.print() // 30->20->10

	myLinkedList.addEnd(node4)
	myLinkedList.addEnd(node5)
	myLinkedList.print()
	// 30->20->10->40->50

	fmt.Println("********")
	myLinkedList.deleteWithValue(10)
	myLinkedList.print()
	// 30->20->40->50
}

// add at front 
func (l *linkedList) addFront(n *node) {
	n.next = l.head
	l.head = n
	l.size++
}

// add at end
func (l * linkedList) addEnd (n *node) {
	temp := l.head
	for  temp.next != nil {
		temp = temp.next
	}
 	temp.next = n
}

// print linkedlist
func (l *linkedList) print() {
	// temp := l.head
	for temp := l.head; temp != nil; {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

// delete with value
func (l *linkedList) deleteWithValue (key int) {
	temp := l.head
	var prev *node
	for temp != nil &&  temp.data != key {
		prev = temp
		temp = temp.next
	}
	if temp != nil {
		prev.next = temp.next
	}
}