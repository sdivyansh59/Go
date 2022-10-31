package main

import "fmt"

// Stack represent stack that hold a slice
type Stack struct {
	 items [] int
}

func main() {
	fmt.Println("Stack in GO")
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)

	fmt.Println(myStack)

	fmt.Printf("poped Val: %v\n",myStack.Pop())
	fmt.Printf("poped Val: %v\n",myStack.Pop())
	fmt.Printf("poped Val: %v\n",myStack.Pop())
	fmt.Printf("poped Val: %v\n",myStack.Pop())
	fmt.Println(myStack)


}

// Push
func (s *Stack) Push (i int){
	s.items = append(s.items, i)
	// remove * and print and see 
	//fmt.Println(s)
}

// Pop will remove ele from the end and
// RETURNs the poped value
func (s *Stack) Pop () int {
	l := len(s.items)-1
	toPop := s.items[l]
	s.items = s.items[:l]

    return toPop
}