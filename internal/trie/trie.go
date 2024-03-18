package trie

import "fmt"

type Trie[T any] struct {
	children map[rune]*Trie[T]
	terminal bool
	value    *T
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		children: make(map[rune]*Trie[T]),
	}
}

func (n *Trie[T]) Put(key string, value T) {
	currentNode := n
	for _, c := range key {
		if _, ok := currentNode.children[c]; !ok {
			currentNode.children[c] = &Trie[T]{
				children: make(map[rune]*Trie[T]),
			}
		}
		currentNode = currentNode.children[c]
	}

	currentNode.value = &value
	currentNode.terminal = true
}

func (n *Trie[T]) Get(key string) (*T, bool) {
	currentNode := n
	for _, c := range key {
		if _, ok := currentNode.children[c]; !ok {
			return nil, false
		}
		currentNode = currentNode.children[c]
	}

	return currentNode.value, true
}

func (n *Trie[T]) Show() {
	currentNode := n
	for k := range currentNode.children {
		fmt.Printf("%s ", string(k))
	}

}

func (n *Trie[T]) Match(key string) (*T, bool) {
	currentNode := n
	for _, c := range key {
		if _, ok := currentNode.children[c]; !ok {
			if currentNode.value != nil {
				return currentNode.value, true
			}

			return nil, false
		}

		currentNode = currentNode.children[c]
	}
	if currentNode.value == nil {
		return nil, false
	}

	return currentNode.value, true
}
