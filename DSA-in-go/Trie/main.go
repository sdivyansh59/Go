package main

import "fmt"


// Article : https://www.educative.io/answers/how-to-create-a-trie-in-go

type TrieNode struct{
	children [26] *TrieNode 
	wordEnd bool
}

// initializing the root of the trie
type Trie struct {
	root * TrieNode
}

// Passing/inserting word to trie
func (t  *Trie) insert(word string) {
	current := t.root

	for _, v := range word {
		idx := v - 'a'
		if current.children[idx] == nil {
			current.children[idx] = new (TrieNode)
		}
		current = current.children[idx]
	}
	current.wordEnd = true
}

//  search for word in trie
func (t * Trie) search(word string) bool{
	current := t.root
	for _,v := range word {
		idx := v -'a'
		if current.children[idx] == nil {
			return false
		}
		current = current.children[idx]
	}
	if current.wordEnd {
		return true
	}
	return false
}

//initializing a new trie
func NewTrie() * Trie {
	t := new (Trie)
	t.root = new (TrieNode)
	return t
}


func main() {
	trie := NewTrie()

	word := [] string{"aqua", "jack", "card"}

	for _,v := range word {
		trie.insert(v)
	}

	words_Search := []string{"aqua", "jack", "card", "care","cat"}

	for _,v := range words_Search{
		if trie.search(v) {
			fmt.Printf(" %s is found in trie \n", v)
		}else {
			fmt.Printf(" %s is not found in trie \n", v)
		}
	}
	
}


