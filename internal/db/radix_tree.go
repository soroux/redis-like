package db

type RadixNode struct {
	children map[byte]*RadixNode
	isEnd    bool
	value    interface{}
}

type RadixTree struct {
	root *RadixNode
}

func NewRadixTree() *RadixTree {
	return &RadixTree{root: &RadixNode{children: make(map[byte]*RadixNode)}}
}

func (t *RadixTree) Insert(key string, value interface{}) {
	node := t.root
	for i := 0; i < len(key); i++ {
		char := key[i]
		if _, ok := node.children[char]; !ok {
			node.children[char] = &RadixNode{children: make(map[byte]*RadixNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
	node.value = value
}

func (t *RadixTree) Get(key string) (interface{}, bool) {
	node := t.root
	for i := 0; i < len(key); i++ {
		char := key[i]
		if _, ok := node.children[char]; !ok {
			return nil, false
		}
		node = node.children[char]
	}
	if node.isEnd {
		return node.value, true
	}
	return nil, false
}
