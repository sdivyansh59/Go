package linkedlist

import (
	"bytes"
	"fmt"
)

type node struct {
	value any
	next *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (l *LinkedList) Insert(newVal any) {
	newNode := &node {
		value: newVal,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) PrintList () {
	temp := l.head
	var strbuff bytes.Buffer

	for temp != nil {
		if strbuff.Len() == 0 {
			strbuff.WriteString(fmt.Sprintf("%v",temp.value))
		}else {
			strbuff.WriteString(fmt.Sprintf("-> %v ",temp.value))
		}
		temp = temp.next
	}

	fmt.Println(strbuff.String()) 
}

func (l *LinkedList) Delete (target any) {
	temp := l.head
	var prev *node 

	for temp != nil {
		if temp.value == target {
			if temp == l.head {
				l.head = l.head.next
			}else {
				prev.next  = temp.next
			}
			return
		}
		prev = temp
		temp = temp.next
	}
	
}

