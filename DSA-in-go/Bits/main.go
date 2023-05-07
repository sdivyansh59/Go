package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 4000000000000000000
	fmt.Println(unsafe.Sizeof(x),x, x>>1)
	fmt.Println(1 & 0)
}