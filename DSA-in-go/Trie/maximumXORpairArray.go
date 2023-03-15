package main

import (
	"fmt"
	"math"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	set   *TrieNode
	unset *TrieNode
}

func (t *Trie) insert(x int) {
	current := t.root

	for i := 31; i >= 0; i-- {
		bit := (x >> i) & 1
		if bit == 1 {
			if current.set == nil {
				current.set = new(TrieNode)
			}
			current = current.set
		} else {
			if current.unset == nil {
				current.unset = new(TrieNode)
			}
			current = current.unset
		}
	}
}

func (t *Trie) findMaxXor(x int) int{
	current := t.root

	var tempMaxXor int
	for i := 31; i >= 0; i-- {

		bit := (x >> i) & 1

		// if both are nil
		if current.set == nil && current.unset == nil {
			break
		}

		if bit == 1 {
			// move to unset bit
			if current.unset == nil {
				current = current.set
				// tempMaxXor += int(math.Pow(float64(2), float64(i)))
			} else {
				current = current.unset
				tempMaxXor += int(math.Pow(float64(2), float64(i)))
			}
		} else {
			//move to set bit
			if current.set != nil {
				current = current.set
				tempMaxXor += int(math.Pow(float64(2), float64(i)))
			} else {
				current = current.unset
				
			}
		}
		
	}
	return tempMaxXor
}

func main() {
	ans := solve([]int{1, 2, 3}, []int{4, 1, 2})
	fmt.Println(ans)
}

func solve(A []int, B []int) int {
	trie := new(Trie)
	trie.root = new(TrieNode)

	for _,v:= range A {
		trie.insert(v)
	}
	maxXor := math.MinInt32

	for _,v := range B {
		smallAns := trie.findMaxXor(v)
		if smallAns > maxXor {
			maxXor = smallAns
		}
	}

	return maxXor
}
