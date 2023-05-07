package main

import (
	"fmt"
	"sync"
)

// tree Node
type TreeNode struct {
	data int
	left *TreeNode
	right *TreeNode
}

type Bst struct {
	root *TreeNode
	lock sync.RWMutex
}

// Insert
func (t *Bst) Insert (val int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.root == nil {
		t.root = &TreeNode{data: val}
	}else {
		t.root = InsertNode(t.root,val)
	}
}
func InsertNode(root *TreeNode,val int) *TreeNode {
	// for any type getting error: not comparable
	if root == nil {
		return &TreeNode{data: val}
	}
	
	if root.data > val {
		root.left = InsertNode(root.left,val)
	}else {
		root.right = InsertNode(root.right,val)
	}
	return root
}


// In-Order
func (t *Bst) InOrderTraversal() {
	t.lock.RLock()
	defer t.lock.RUnlock()
	fmt.Print("In-Order: ")
	inOrderTraversal(t.root)
	fmt.Println("")
}
func inOrderTraversal(root *TreeNode){
	if root == nil {
		return 
	}
	inOrderTraversal(root.left)
	fmt.Print(root.data ," ")
	inOrderTraversal(root.right)
}


//Pre-Order
func (t *Bst) PreOrderTraversal () {
	t.lock.RLock()
	defer t.lock.RUnlock()
	fmt.Print("Pre-Order: ")
	preOrderTraversal(t.root)
	fmt.Println("")
}
func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return 
	}
	fmt.Print(root.data, " ")
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
}


// Post-Order
func(t *Bst) PostOrderTraversal() {
	t.lock.RLock()
	defer t.lock.RUnlock()
	fmt.Print("Post-Order: ")
	postOrderTraversal(t.root)
	fmt.Println("")
}
func postOrderTraversal(root *TreeNode) {
	if root == nil {
		return 
	}
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
	fmt.Print(root.data, " ")
}


// Binary Search
func (t *Bst) SearchNode(target int) bool{
	t.lock.RLock()
	defer t.lock.RUnlock()

    if searchNode(t.root, target) {
		fmt.Printf("%v is present\n", target)
		return true
	}else {
		fmt.Printf("%v is not present\n", target)
		return false
	}
}
func searchNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.data > target { 
		// move left
		return searchNode(root.left,target)
	}
	if root.data < target {
		// mover right
		return searchNode(root.right, target)
	}
	// Got it: root.data == target
	return true
}


//Delete->Time Complexity: O(h) | Auxilary space(recursive stack): O(h)
// where h = height of tree (in case of skewed tree h==n)
func (t *Bst) DeleteNode (targetVal int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	deleteNode(t.root, targetVal)
}
func deleteNode(root *TreeNode, target int) *TreeNode{
	if root == nil {
		return nil
	}

	if root.data > target {
		root.left = deleteNode(root.left, target)
	}else if root.data < target {
		root.right = deleteNode(root.right, target)
	}else { // we are at target TreeNode
		// one child case
		if root.left == nil {
			return root.right
		}else if root.right == nil {
			return root.left
		}else { // two child case
			var successor *TreeNode = getSucc(root)
			root.data = successor.data
			root.right = deleteNode(root.right, successor.data)
		}
	}
	return root
}
//Note:logic of finding successor works here but this is not 
//general  approach to find successor of node.
func getSucc(root *TreeNode) *TreeNode{
	// subtree ka left most ele find karna hai
	curr := root.right
	for curr != nil && curr.left != nil {
		curr = curr.left
	}
	return curr
}


func main() {
	fmt.Println("BST implementation")

	bst := Bst{}

	bst.Insert(200)
	bst.Insert(100)
	bst.Insert(400)
	bst.Insert(300)
	bst.Insert(50)
	
	bst.InOrderTraversal()
	// bst.PostOrderTraversal()
	// bst.PreOrderTraversal()

	bst.SearchNode(300)
	bst.SearchNode(500)

	bst.DeleteNode(400)
	bst.InOrderTraversal()

}





