package genver

import (
	"time"

	"github.com/gregfurman/genver/internal/build"
	"github.com/gregfurman/genver/pkg/store"
)

var (
	buildInfo build.BuildInfo
)

func init() {
	buildInfo = build.GetBuildInfoFromRuntime()
}

func NewDependencyVersionStore() *store.DependencyStore {
	return store.NewStore(buildInfo.Dependencies)
}

func GetVersion() string {
	return buildInfo.VersionInfo.Version
}

func GetRevision() string {
	return buildInfo.VersionInfo.Revision
}

func GetLastCommit() time.Time {
	return buildInfo.VersionInfo.LastCommit
}

func GetIsDirtyBuild() bool {
	return buildInfo.VersionInfo.DirtyBuild
}
