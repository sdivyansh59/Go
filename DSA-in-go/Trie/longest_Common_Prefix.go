package main

import (
	"fmt"
)

type TrieNode struct {
	children [26]*TrieNode
	wordEnd  bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) lcp(word string) [][]string {
	current := t.root
	var lcpStr string
	ans := make([][]string, 0)

	for _, v := range word {
		idx := v - 'a'

		if current.children[idx] == nil {
			// return lcpStr
			break
		} else {
			smallAns := make([]string, 0)
			lcpStr += string(v)
			findAllPrefix(current.children[idx], lcpStr, &smallAns)
			if len(smallAns) > 0 {
				ans = append(ans, smallAns)

			}
		}

		current = current.children[idx]
	}

	return ans
}

func findAllPrefix(current *TrieNode, lcpStr string, smallOutput *[]string) {

	if len(*smallOutput) >= 3 {
		return
	}

	if current.wordEnd {
		*smallOutput = append(*smallOutput, lcpStr)
		// return
	}

	for i, v := range current.children {
		if v != nil {
			// lcpStr += string('a' + i)
			findAllPrefix(v, lcpStr+string('a' + i), smallOutput)
		}
	}
}

func (t *Trie) insert(word string) {
	current := t.root

	for _, v := range word {
		idx := v - 'a'
		if current.children[idx] == nil {
			current.children[idx] = new(TrieNode)
		}
		current = current.children[idx]
	}
	current.wordEnd = true
}

func getProductSuggestions(products []string, search string) [][]string {
	// Write your code here
	trie := new(Trie)
	trie.root = new(TrieNode)

	for _, v := range products {
		trie.insert(v)
	}

	ans := trie.lcp(search)

	return ans

}

func main() {
	str := []string{"carpet", "cart", "car", "camera", "crate"}

	ans := getProductSuggestions(str, "camera")
	fmt.Println(ans)
}