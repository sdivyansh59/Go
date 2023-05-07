package main

import "fmt"

const mod int = 1000000007

func main() {
	l := []int{1, 2, 3, 4, 5}
	for i := range l {
		fmt.Println(i)
	}
}



func levelOrder(A *treeNode )  ([][]int) {
    
    queue := [] *treeNode{}
    queue = append(queue,A)
    
    llist := [][]int{}
    for i,_ :=  range llist{
        llist[i] = []int{}
    }
    
    // lcount := 0
    for len(queue) > 0 {
        k := len(queue) 
        thislevel := []int {}
        for k >0 {
            t := queue[0]
            queue = queue[1:]
            thislevel = append(thislevel,t)
            k--
            
            if t.left != nil {
                queue = append(queue,t.left)
            }
            if t.right != nil {
                queue = append(queue,t.right)
            }
        }
        // next level
        llist =append(llist,thislevel)
    }
    return llist
}
