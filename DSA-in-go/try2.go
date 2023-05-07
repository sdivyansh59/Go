package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello")
	
}

func print( head *listNode) {
	var str string
	for head != nil {
		str += fmt.Sprintf("%v ",head.value)
		head = head.next
	}
	fmt.Println(str)
}

func solve(A *listNode )  (*listNode) {
	if A == nil || A.next == nil {
		return A
	}

	odd,even := A, A.next
	var evenHead, t2 *listNode

	for even != nil {
		odd.next = even.next

		if evenHead == nil {
			evenHead = even
			t2 = even
		}else {
			t2.next = even
			t2 = t2.next
		}

		odd = even.next
		if odd != nil {
			even = odd.next
		}else {
			break
		}
	}

	t2.next = nil

	print(A)
	print(evenHead)

	evenHead = reverse(evenHead)

	insertEvenNode(A, evenHead)
	return A


}


func reverse( head *listNode) *listNode {
	if head == nil || head.next == nil {
		return head
	}

	var next *listNode

	first , second := head, head.next

	for second != nil {
		next = second.next
		second.next = first

		first = second
		second = next
	}

	head.next = nil

   return first

}

func insertEvenNode(oddHead , evenHead *listNode) *listNode {
	head := oddHead
	
	for evenHead != nil {
		next = evenHead.next

		evenHead.next = oddHead.next
		oddHead.next = evenHead

		oddHead = evenHead.next
		evenHead = next
		 
	}

	return head
}


