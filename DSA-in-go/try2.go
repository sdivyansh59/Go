package main

import "sort"

func main() {

}


func solve(A []int , B int )  ([]int) {

	for i,v := range A {
		idx := -1
        if B == 0 {
			break
		}
		for j:= i+1 ; j< len(A) ; j++ {
			if A[j] > v {
				if idx != -1 {
					if A[idx] < A[j] {
						idx = j
					}
				}else {
					idx = j
					B--
				}
			}
		}

		A[i], A[idx]= A[idx], A[i]


	}
	return A
}