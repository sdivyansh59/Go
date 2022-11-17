package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return &Node{data: val}
	}
	if n.data < val { // move right
		n.right = n.right.Insert(val)

	} else if n.data > val {
		n.left = n.left.Insert(val)
	}
	// if equal, don't insert anything
	return n
}

func (n *Node) PrintBST() {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	n.left.PrintBST()
	n.right.PrintBST()
}

func main() {
	var root *Node
	// fmt.Println(root)

	root = root.Insert(10)
	// fmt.Println(root)
	root.Insert(15)
	root.Insert(8)
	root.Insert(12)
	root.Insert(20)
	root.Insert(5)

	root.PrintBST()

}

func (root *Node) levelOrder() {

}