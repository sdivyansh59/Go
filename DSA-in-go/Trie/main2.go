package main

type TrieNode struct {
	children [26]*TrieNode
	wordEnd  bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	trie := new(Trie)
	trie.root = new(TrieNode)
	return *trie
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

func (t *Trie) search(word string) bool {
	current := t.root
	for _, v := range word {
		idx := v - 'a'
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

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root

	for _, v := range prefix {
		idx := v - 'a'
		if current.children[idx] == nil {
			return false
		}

	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */