package leetcode

// 使用可变长度的前缀树
type trieNode struct {
	end  bool
	next map[rune]*trieNode
}

type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{
		root: &trieNode{
			next: map[rune]*trieNode{},
		},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, r := range word {
		if v, ok := node.next[r]; ok {
			// 存在，继续
			node = v
		} else {
			node.next[r] = &trieNode{
				next: map[rune]*trieNode{},
			}
			node = node.next[r]
		}
	}
	node.end = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, r := range word {
		if v, ok := node.next[r]; ok {
			node = v
		} else {
			return false
		}
	}
	// 未结束需要返回false
	if !node.end {
		return false
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, r := range prefix {
		if v, ok := node.next[r]; ok {
			node = v
		} else {
			return false
		}
	}
	return true
}
