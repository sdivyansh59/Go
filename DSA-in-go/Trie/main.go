package main

type TrieNode struct{
	children [26] *TrieNode 
	wordEnd bool
}

// initializing the root of the trie
type Trie struct {
	root * TrieNode
}

//initializing a new trie
func NewTrie() * Trie {
	t := new (Trie)
	t.root = new (TrieNode)
	return t
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


func main() {
	trie := NewTrie()

	word := [] {"aqua", "jack", "card"}

}
