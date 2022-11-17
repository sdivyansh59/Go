package main

import (
	"fmt"
	"sort"
)

func main() {
	mymap := map[int] string{}
	mymap[10] = "abhisheak"
	mymap[-2] = "divyansh"
	mymap[100] = "babu"
	mymap[0] = "aakash"

//it will print in sorted order
	fmt.Println(mymap)

// store key of map in it
	keys := []int {}

// not maintain any order
	for k := range mymap {
		keys = append(keys, k)
	}

// sort map A/c to key
	sort.Ints(keys)

	for _,k := range keys {
		fmt.Println(k , mymap[k])
	}

	fmt.Println("*****************")
/* sort map A/C to value of map
we will sort keys slice a/c val of map
*/ // i, j are index
	sort.SliceStable(keys, func (i, j int) bool {
		return mymap[keys[i]] < mymap[keys[j]]	
	})

// print map in sorted order
    for _,k := range keys {
		fmt.Println(k , mymap[k])
	}
}