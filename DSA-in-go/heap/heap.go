package main

import "fmt"

//https://www.youtube.com/watch?v=3DYIgTC4T1o&t=464s&ab_channel=JunminLee


type MaxHeap struct {
	array [] int // actually it is slice
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

// Extract returns the largest key
// remove it from the heap
func (h *MaxHeap) Extract() int  {
	extracted := h.array[0]
	l := len(h.array) -1
	// when the array is empty
	if len(h.array) ==0 {
		fmt.Println("can not extract because array length is 0")
		return -1
	}
	// take out the last idx & put it in the 
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap)maxHeapifyDown(index int) {
	lastIndex := len(h.array)
	l, r := left(index) , right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		// when the left child is the only child
		if l == lastIndex {

		}else if h.array[l] > h.array[r] {
		// when left  child is larger
			childToCompare = l	
		}else {
		// when right child is larger	
			childToCompare = r
		}
	}

}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// get the parent idx
func parent( i int) int {
	return (i-1)/2
}

// get the left child idx
func left (i int) int {
	return 2*i +1
}
// get right child idx
func right (i int) int {
	return 2*i +2
}

// swap keys in the array
func(h *MaxHeap) swap (i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	fmt.Println("we will cover max heap")
	
}