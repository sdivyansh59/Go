package main

import (
	"fmt"
	"sort"
)

type void struct{}

var nothing void

func main() {
	//creating set in go
	set := make(map[string]void)
	set["apple"] = nothing
	set["banana"] = nothing
	set["apple"] = nothing
	set["tomato"] = nothing
	set["potato"] = nothing

	delete(set,"apple")
	fmt.Println(set)
	//print set
	for i,v := range set {
		fmt.Println(i,v)
	}

	// check if element present or not 
	if _,ok := set["tomato"] ; ok {
		fmt.Print(set["golu"])
	}

	slice1 := []int {20,30,-5,1,2}
	fmt.Println(slice1)

	sort.IntSlice.Sort(slice1,(i int ,j int)bool{
		return i <j
	})
	fmt.Println(slice1)


}
