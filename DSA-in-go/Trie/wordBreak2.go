package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children [26] *TrieNode
	wordEnd  bool
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
	return current.wordEnd
}

func wordBreak(s string, wordDict []string) []string {
	trie := new(Trie)
	trie.root = new(TrieNode)

	for _, v := range wordDict {
		trie.insert(v)
	}

	ans := make([]string,0)

	trie.solve(s, "", &ans)

	for i,j := 0, len(ans)-1 ; i <j ; i,j = i+1, j-1 {
		ans[i] = strings.Trim(ans[i]," ")
		ans[j] = strings.Trim(ans[j]," ")
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans

}

func (t* Trie) solve(s string, output string, ans *[]string) {
	if len(s) == 0 {
		*ans = append(*ans, output)
		fmt.Println(output)
		return
	}
	currStr := ""
	for i,v := range s {
		currStr += string(v)
		if t.search(currStr) {
			t.solve(s[i+1:],output+currStr+" ", ans)
		}
	}
}


func main() {
	ans := wordBreak("catsanddog",[]string {"cat","cats","and","sand","dog"})
	fmt.Println(ans)
}

func Search2(word string, root *TrieNode) bool {
	if len(word) == 0 {
		return root.wordEnd 
	}

	for i,v := range word {
		if v == '.' {
			var smallAns bool
			for j:= 0 ; j<26 ; j++ {
                if root.children[j] != nil {
					smallAns = Search2(word[i+1:],root.children[j])
                    if smallAns {
						break
					}
                    
                }
                
            }
			return smallAns
		}

		idx := v -'a'
		if root.children[idx] == nil {
			return false
		}
		return Search2(word[1:],root.children[idx])
	}


}

func (t * WordDictionary) Search(word string) bool{
	return Search2(word,t.root)
}
