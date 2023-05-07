package main

type listNode struct {
	value int
	next  *listNode
}

func findMid(start , end *listNode) *listNode{
	var slow, fast *listNode = start, start

	for (fast != nil && fast != end) && (fast.next != nil && fast.next != end) {
		slow = slow.next
		fast = slow.next
	}

	return slow
}

func merge (head1 , end1, head2, end2 *listNode){
	var Head , head *listNode
	var temp1, temp2 *listNode = head1, head2
	
	for temp1 != 

}


func sortList(A *listNode) *listNode {

}
