package main

import "fmt"

/* Article on Set
https://linuxhint.com/golang-set/#:~:text=A%20set%20refers%20to%20a,structs%20to%20create%20a%20set.
*/

// it take 0kb memory
type void struct {}
var nothing void 

func main () {
	set := make(map[string]void)
	set["apple"] = nothing
	set["orange"] = nothing
	set["mango"] = nothing
	fmt.Println(set)

	// iterate set
	for k := range set {
		fmt.Println(k)
	}

	// Delete element from set
	delete(set, "mango")
	for k := range set {
		fmt.Println("after delete: ", k)
	}

	// chek if element exist in set
	if _, ok := set["apple"]; ok {
		fmt.Println("element exist")
	}else {
		fmt.Println("element not exist")
	}
}