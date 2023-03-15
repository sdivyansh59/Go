package main

import "container/heap"

// type void struct{}

func main() {

	
	
	 

}


type IntHeap []int // mn heap

func (h IntHeap) Len() int { return len(h)}
func (h IntHeap) Less(i,j int) bool {return h[i] < h[j]}
func (h IntHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func maxp3(A []int )  (int) {
    h := &IntHeap{}
	heap.Init(h)

	for i := 0; i< len(A) ; i++ {
		if A[i] !=0 {
			if h.Len() >0 && (*h)[0] < A[i] {
				if h.Len() >3 {
					heap.Pop(h)
				}	
				heap.Push(h,A[i])
			}-*
		}
	}

	output := 1
	for h.Len() >0 {
		output = output* heap.Pop(h).(int)
	}
	return output
}
