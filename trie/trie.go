package trie

import (
	"sync"
)

// Trie Tree
type Trie struct {
	Root           *Node
	Mutex          sync.RWMutex
	//CheckWhiteList bool // 是否检查白名单
}

// Node Trie tree node
type Node struct {
	Node map[rune]*Node
	End  bool
}

// NewTrie returns a Trie tree
func NewTrie() *Trie {
	t := new(Trie)
	t.Root = NewTrieNode()
	return t
}

// NewTrieNode return a *TrieNode
func NewTrieNode() *Node {
	n := new(Node)
	n.Node = make(map[rune]*Node)
	n.End = false
	return n
}

// Add 添加一个敏感词(UTF-8的)到Trie树中
func (t *Trie) Add(keyword string) {
	chars := []rune(keyword)

	if len(chars) == 0 {
		return
	}

	t.Mutex.Lock()

	node := t.Root
	for _, char := range chars {
		if _, ok := node.Node[char]; !ok {
			node.Node[char] = NewTrieNode()
		}
		node = node.Node[char]
	}
	node.End = true

	t.Mutex.Unlock()
}

// Query 查询敏感词
// 将text中在trie里的敏感字，替换为*号
// 返回结果: 是否有敏感字, 敏感字数组, 替换后的文本
func (t *Trie) Query(text string) bool {
	var found = 0
	chars := []rune(text)
	l := len(chars)
	if l == 0 {
		return false
	}

	var (
		i, j, jj int
		ok       bool
	)

	node := t.Root
	for i = 0; i < l; i++ {
		if _, ok = node.Node[chars[i]]; !ok {
			continue
		}

		jj = 0

		node = node.Node[chars[i]]
		for j = i + 1; j < l; j++ {
			if _, ok = node.Node[chars[j]]; !ok {
				if jj > 0 {
						found = j+1-i
						i = jj
				}
				break
			}

			node = node.Node[chars[j]]
			if node.End {
				jj = j //还有子节点的情况, 记住上次找到的位置, 以匹配最大串 (eg: AV, AV女优)

				if len(node.Node) == 0 || j+1 == l { // 是最后节点或者最后一个字符, break
						found = j+1-i
						i = j
						break
				}
			}
		}
		node = t.Root
	}

	exist := false
	if found > 0 {
		exist = true
	}
	return exist
}


