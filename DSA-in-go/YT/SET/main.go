package main

import "fmt"

// it take  0kb memory
type void struct{}
var nothing void


func main() {

	set := make(map[any] void)

	// create set of string data type
	set["apple"] = nothing
	set["apple"] = nothing
	set["mango"] = nothing
	set["orange"] = nothing

	fmt.Println(set)


	//How to  iterate over set ?
	for k := range set {
		fmt.Println(k)
	}
	// insertion order is not maintain in map/set

	// How delete element from set
	delete(set,"mango")
	fmt.Println(set)

	// check if element exist in set
	if _, ok := set["mango"] ; ok {
		fmt.Println("mango is there in my set")
	}else {
		fmt.Println("mango is not  in my set")
	}


}