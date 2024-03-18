package trie

import (
	"strings"
)

type Trie[T any] struct {
	children map[string]*Trie[T]
	terminal bool
	value    *T
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		children: make(map[string]*Trie[T]),
	}
}

func (n *Trie[T]) Put(key string, value T) {
	parts := strings.FieldsFunc(key, func(r rune) bool { return r == '/' || r == '-' || r == '.' || r == '_' })

	currentNode := n
	for _, part := range parts {
		if _, ok := currentNode.children[part]; !ok {
			currentNode.children[part] = &Trie[T]{
				children: make(map[string]*Trie[T]),
			}
		}
		currentNode = currentNode.children[part]
	}

	currentNode.value = &value
	currentNode.terminal = true
}

func (n *Trie[T]) Get(key string) (*T, bool) {
	parts := strings.FieldsFunc(key, func(r rune) bool { return r == '/' || r == '-' || r == '.' || r == '_' })

	currentNode := n
	for _, part := range parts {
		if _, ok := currentNode.children[part]; !ok {
			return nil, false
		}
		currentNode = currentNode.children[part]
	}

	return currentNode.value, true
}

func (n *Trie[T]) Match(key string) (*T, bool) {
	parts := strings.FieldsFunc(key, func(r rune) bool { return r == '/' || r == '-' || r == '.' || r == '_' })

	currentNode := n
	for _, part := range parts {
		if _, ok := currentNode.children[part]; !ok {
			if currentNode.value != nil {
				return currentNode.value, true
			}

			return nil, false
		}

		currentNode = currentNode.children[part]
	}
	if currentNode.value == nil {
		return nil, false
	}

	return currentNode.value, true
}
