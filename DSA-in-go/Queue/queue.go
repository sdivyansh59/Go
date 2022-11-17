package main

import "fmt"

// Queue represent a queu that hold a slice
type Queue struct {
	items [] string
}

func main() {
	q := Queue{}
	fmt.Println(q)

	q.Enque("Divyansh")
	q.Enque("Abhisheak")
	q.Enque("AAkash")
	fmt.Println(q)

	fmt.Println(q.Deqeue())
	fmt.Println(q)
}

// Enque add element at end
func (q *Queue) Enque(newVal string) {
	q.items = append(q.items, newVal)
}

// Deqeue remove ele from front
func (q *Queue) Deqeue() string{
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

// isEmpty
func (q *Queue) isEmpty() bool{
	if len(q.items) == 0 {
		return true
	}else {
		return false
	}
}