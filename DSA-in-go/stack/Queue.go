package main

import "fmt"

//Stack represent the stack that hold the slice
type Queue struct {
	items [] int
}

func main() {
	fmt.Println("Queue and Dequeue")
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Push(10)
	myQueue.Push(20)
	myQueue.Push(30)
	fmt.Println(myQueue)

	myQueue.Pop()
	fmt.Println(myQueue)



}

// Push will add vlue at the end and return it
func (q *Queue) Push (i int) int {
	q.items = append(q.items, i)
	return i
}

//Pop will remove the value from the front and
// RETURNs the poped value
func (q *Queue) Pop () int {
	if len(q.items) <=0 {
		return -1
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}