package main

 type treeNode struct {
 *     left *treeNode
 *     value int
 *     right *treeNode
 * }
 * 
 * func treeNode_new(val int) *treeNode{
 *     var node *treeNode = new(treeNode)
 *     node.value=val
 *     node.left=nil
 *     node.right=nil
 *     return node
 * }

 

func main() {

}

func solve(root *treeNode) []int {
	m := map[int] int{}
    levelOrder(root, m)
    
    keys := []int
    for k,v := range m {
        keys = append(keys,k)
    }
    
    sort.Ints(keys)
	fmt.Println(keys)

	rightView := [] int
	for _, val := range keys{
		rightView = append(rightView, m[val])
	}  

	return rightView
}




