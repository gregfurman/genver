package trie_test

import (
	"testing"

	"github.com/gregfurman/genver/internal/build"
	"github.com/gregfurman/genver/internal/trie"
)

var buildDeps []build.DepInfo = []build.DepInfo{
	{Path: "github.com/aws/aws-sdk-go-v2", Version: "v1.25.2"},
	{Path: "github.com/aws/aws-sdk-go", Version: "v1.25.1"},
	{Path: "github.com/aws/smithy-go", Version: "v1.20.1"},
	{Path: "github.com/davecgh/go-spew", Version: "v1.1.1"},
	{Path: "github.com/google/go-cmp", Version: "v0.5.8"},
	{Path: "github.com/jmespath/go-jmespath", Version: "v0.4.0"},
	{Path: "github.com/pmezard/go-difflib", Version: "v1.0.0"},
	{Path: "github.com/sirupsen/logrus", Version: "v1.9.3"},
	{Path: "github.com/stretchr/objx", Version: "v0.1.0"},
	{Path: "github.com/stretchr/testify", Version: "v1.7.0"},
	{Path: "golang.org/x/mod", Version: "v0.8.0"},
	{Path: "golang.org/x/sys", Version: "v0.5.0"},
	{Path: "golang.org/x/text", Version: "v0.14.0"},
	{Path: "golang.org/x/tools", Version: "v0.6.0"},
	{Path: "gopkg.in/check.v1", Version: "v0.0.0-20161208181325-20d25e280405"},
	{Path: "gopkg.in/yaml.v3", Version: "v3.0.0-20200313102051-9f266ea9e77c"},
	{Path: "gitlab.com/company/repos/ecr/promoter.git", Version: "v0.1.0"},
}

func Benchmark_Store(b *testing.B) {
	for n := 0; n < b.N; n++ {
		t := trie.NewTrie[string]()
		for _, dep := range buildDeps {
			t.Put(dep.Path, dep.Version)
		}
	}
}

func Benchmark_Match(b *testing.B) {
	t := trie.NewTrie[string]()
	for _, dep := range buildDeps {
		t.Put(dep.Path, dep.Version)
	}
	for n := 0; n < b.N; n++ {
		for _, dep := range buildDeps {
			t.Match(dep.Path)
		}
	}
}
