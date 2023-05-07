package main

type listNode struct {
	value int
	next  *listNode
}

// func solve(A *listNode )  (*listNode) {

// 	k

// }

func rotateRight(A *listNode, B int) *listNode {
	var len int = 1
	var tail, temp, head *listNode = nil, A, nil
	
	for temp.next != nil {
		len ++
		temp = temp.next
	}

	tail = temp
	B= ( len %B)
	len = len-B
    
	temp = A
	for len >1 {
		len--
		temp = temp.next
	}

	head = temp.next
	temp.next = nil
	tail.next = A

	return head
}
