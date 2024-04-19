package store

import (
	"reflect"

	"github.com/gregfurman/genver/internal/build"
	"github.com/gregfurman/genver/internal/trie"
)

type DependencyStore struct {
	dependencyTrie *trie.Trie[string]
}

func NewStore(deps []build.DepInfo) *DependencyStore {
	depTrie := trie.NewTrie[string]()
	for _, dep := range deps {
		depTrie.Put(dep.Path, dep.Version)
	}

	return &DependencyStore{
		dependencyTrie: depTrie,
	}
}

func (s *DependencyStore) FindVersionFromPath(path string) any {
	value, ok := s.dependencyTrie.PrefixMatch(path)
	if ok {
		return *value
	}

	return nil
}

func (s *DependencyStore) FindVersionFromData(data any) any {
	value, ok := s.dependencyTrie.PrefixMatch(reflect.TypeOf(data).PkgPath())
	if ok {
		return *value
	}

	return nil
}
