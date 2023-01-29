package main

import (
	"fmt"
	"sort"
	"strings"
)

type student struct {
	name string
	id int
}

func main() {
	list := []student{
		{"Divyansh ", 2},
		{"Sonam ", 20},
		{"AAlam ", 3},
	}
    sort.Slice(list, func(i, j int) bool {
		//return list[i].id < list[j].id
		if strings.Compare(list[i].name ,list[j].name) < 0 {
			return true
		}
		return false
	})

	arr := [] string {"apple","zoy","ball"}
	sort.Strings(arr)

	fmt.Println(list, arr)
}