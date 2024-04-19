package store_test

import (
	"testing"

	"github.com/gregfurman/genver/internal/build"
	"github.com/gregfurman/genver/pkg/store"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Test_Store_FindVersionFromPath(t *testing.T) {
	deps := []build.DepInfo{
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

	modulePathsExtensions := []string{
		"/pkg", "/pkg/service", "/pkg/service/generate_version",
	}

	s := store.NewStore(deps)
	for _, want := range deps {
		if got := s.FindVersionFromPath(want.Path); got != want.Version {
			t.Errorf("retrieved incorrect version for dependency '%s'. Got '%v', want '%s'", want.Path, got, want.Version)
		}

		for _, m := range modulePathsExtensions {
			path := want.Path + m
			if got := s.FindVersionFromPath(path); got != want.Version {
				t.Errorf("retrieved incorrect version for dependency '%s'. Got '%v', want '%s'", path, got, want.Version)
			}
		}
	}

	noExistDeps := []build.DepInfo{
		{Path: "gopkg.in/json.v3", Version: "v3.0.0-20200313102051-9f266ea9e77c"},
		{Path: "github.com/ai/helicopter", Version: "v9.0.1"},
	}
	for _, want := range noExistDeps {
		if got := s.FindVersionFromPath(want.Path); got != nil {
			t.Errorf("retrieved incorrect version for dependency '%s'. Got '%v', want '%s'", want.Path, got, "")
		}
	}
}

func Test_Store_FindVersionFromData(t *testing.T) {
	deps := []build.DepInfo{
		{Path: "golang.org/x/text", Version: "v0.14.0"},
	}

	tests := []struct {
		pkgName string
		arg     any
		want    string
	}{
		{pkgName: "golang.org/x/text/cases", arg: cases.Caser{}, want: "v0.14.0"},
		{pkgName: "golang.org/x/text/language", arg: language.English, want: "v0.14.0"},
	}

	s := store.NewStore(deps)
	for _, test := range tests {
		if got := s.FindVersionFromData(test.arg); got != test.want {
			t.Errorf("retrieved incorrect version for dependency '%s'. Got '%v', want '%s'", test.pkgName, got, test.want)
		}
	}
}
