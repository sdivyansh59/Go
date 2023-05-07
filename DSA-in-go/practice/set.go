package main

import "fmt"

type Stack struct {
	items []any
}

func (s *Stack) Push (x any) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}


func main() {
	st := Stack{}

	st.Push("Divyansh")
	st.Push("Ankit")
	st.Push("Priyanka")
	
	st.Pop()
	fmt.Println(st)

}