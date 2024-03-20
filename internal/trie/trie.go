package trie

import (
	"strings"
)

type Trie[T any] struct {
	children map[string]*Trie[T]
	value    *T
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		children: make(map[string]*Trie[T]),
	}
}

// Put inserts a value into the trie corresponding to a given key.
// The key is broken up into a set of traversable parts, with each part
// being stored as a node and the value being stored in the terminal node.
func (n *Trie[T]) Put(key string, value T) {
	parts := strings.Split(key, "/")

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
}

// Get searches for an exact match between an input key and
// a chain of nodes in the trie. Returns a pointer to the associated value
// and a boolean indicating whether an exact match was found.
func (n *Trie[T]) Get(key string) (*T, bool) {
	parts := strings.Split(key, "/")

	currentNode := n
	for _, part := range parts {
		if _, ok := currentNode.children[part]; !ok {
			return nil, false
		}
		currentNode = currentNode.children[part]
	}

	return currentNode.value, true
}

// PrefixMatch searches for the longest key prefix that matches the given key string.
// It returns a pointer to the value associated with the longest matching prefix key,
// and a boolean indicating whether a match was found.
func (n *Trie[T]) PrefixMatch(key string) (*T, bool) {
	parts := strings.Split(key, "/")

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
