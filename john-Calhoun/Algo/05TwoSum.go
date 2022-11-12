package main

import (
	"fmt"
	"sort"
)

func twoSum(list []int, key int) (int, int) {
	// sort the list then use 2 pointer approach
	sort.Ints(list)
	fmt.Println(list)

	i := 0
	j:= len(list)-1
	for i <j {
		if list[i] +list[j] == key {
			return i,j
		}else if list[i] +list[j] < key {
			i++
		}else {
			j--
		}
	}
	return -1,-1
}

func main() {
	list := [] int {1,2,3,4,5,100,-2}
    key := 0
	idx1 , idx2 := twoSum(list, key)
	fmt.Println(idx1, idx2)
	
}