package build

import (
	"encoding/json"
	"net/url"
	"runtime/debug"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type BuildInfo struct {
	VersionInfo  VerInfo
	Dependencies []DepInfo
}

type VerInfo struct {
	Version    string
	Revision   string
	LastCommit time.Time
	DirtyBuild bool
}

type DepInfo struct {
	Version    string `json:"Version"`
	Path       string `json:"Path"`
	Name       string `json:"Name"`
	Main       bool   `json:"Main"`
	IsIndirect bool   `json:"Indirect"`
}

func GetBuildInfoFromRuntime() BuildInfo {
	info, _ := debug.ReadBuildInfo()
	return BuildInfo{
		VersionInfo:  getVersionInfo(info.Main.Version, info.Settings),
		Dependencies: getDependencyInfo(info.Deps),
	}
}

func GetBuildInfoFromData(content string) BuildInfo {
	info, _ := debug.ReadBuildInfo()
	return BuildInfo{
		VersionInfo:  getVersionInfo(info.Main.Version, info.Settings),
		Dependencies: getDependencyInfoFromData(content),
	}
}

func getVersionInfo(version string, settings []debug.BuildSetting) VerInfo {
	var vers VerInfo = VerInfo{Version: version}
	for _, setting := range settings {
		if setting.Value == "" {
			continue
		}
		switch setting.Key {
		case "vcs.revision":
			vers.Revision = setting.Value
		case "vcs.time":
			vers.LastCommit, _ = time.Parse(time.RFC3339, setting.Value)
		case "vcs.modified":
			vers.DirtyBuild = setting.Value == "true"
		}
	}

	return vers
}

func getDependencyInfo(dependencies []*debug.Module) []DepInfo {
	var deps []DepInfo
	for _, dep := range dependencies {
		if dep.Sum != "" {
			deps = append(deps, DepInfo{Version: dep.Version, Path: dep.Path, Name: nameFromPath(dep.Path)})
		}
	}

	return deps
}

func getDependencyInfoFromData(content string) []DepInfo {
	var deps []DepInfo

	if err := json.Unmarshal([]byte(fmtJson(content)), &deps); err != nil {
		return nil
	}

	var filteredDeps []DepInfo
	for _, dep := range deps {
		if dep.Main || dep.Version == "" || dep.Path == "" {
			continue
		}
		filteredDeps = append(filteredDeps, DepInfo{Version: dep.Version, Path: dep.Path, Name: nameFromPath(dep.Path)})
	}

	return filteredDeps
}

func nameFromPath(p string) string {
	u, err := url.Parse("//" + p)
	if u == nil || err != nil {
		return ""
	}

	// split the pathname on the following runes
	parts := strings.FieldsFunc(u.Path, func(r rune) bool { return r == '/' || r == '-' || r == '.' || r == '_' })

	var sb strings.Builder
	caser := cases.Title(language.Und, cases.NoLower)

	// convert the package path to camelcase i.e "foo.bar.golang.org/pkg/service" -> "PkgService"
	for _, part := range parts {
		sb.WriteString(caser.String(part))
	}

	// extract the eTLD+1 for a given package hostname i.e eTLD+1 for "foo.bar.golang.org" is "golang.org".
	tldPlusOne, _ := publicsuffix.EffectiveTLDPlusOne(u.Host)

	// just get the domain, no TLD attached i.e "golang.org" -> "golang"
	left, _, ok := strings.Cut(tldPlusOne, ".")

	// if the package has paths AND a valid domain is extracted add a "_{domain}" i.e
	if len(parts) > 0 && ok {
		sb.WriteString("_" + caser.String(left))
	}

	return sb.String()
}

func fmtJson(content string) string {
	collapsed := strings.ReplaceAll(content, "\n", "")
	corrected := strings.ReplaceAll(collapsed, "}{", "},{")

	return "[" + corrected + "]"
}
