package genver

import (
	"github.com/gregfurman/genver/internal/build"
	"github.com/gregfurman/genver/pkg/store"
)

func NewDependencyVersionStore() *store.DependencyStore {
	return store.NewStore(build.GetBuildInfoFromRuntime())
}
