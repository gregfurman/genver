# GenVer

GenVer is a lightweight module that improves accessibility to Go dependencies information.

## Installation

CLI:
```shell
go install github.com/gregfurman/genver/cmd/cli@latest
```

Package:
```shell
go get github.com/gregfurman/genver
```

## Usage

GenVer comes in two flavours --  a CLI and importable package.

### Command-line Interface
The binary takes, as an argument, the full content of your Go modules in JSON form. This can easily be retrieved with `cmd/go` with the `go list -m -json all` command. See [Flags](#flags) for more options and CLI flags.

The generated code file will contain information of each package, defined as two `const` values per module imported, one for module's `Version` and `Path`, respectively. The generic naming of each constant provides information in the form `<PATH>_<DOMAIN>_<TYPE>`. i.e  will be displayed as:


<details open>
<summary>Example dependencies and their generated output</summary>
<br>

#### Example #1 `cloud.google.com/go/analytics@v0.12.0`:
```golang
// Currently using cloud.google.com/go/analytics v0.12.0 
const (
    GoAnalytics_Google_Version = "v0.12.0"
    GoAnalytics_Google_Path    = "cloud.google.com/go/analytics"
)
```

#### Example #2 `go.opencensus.io@v0.24.0`:
```golang
// Currently using go.opencensus.io v0.24.0 
const (
    Opencensus_Version = "v0.24.0"
    Opencensus_Path    = "go.opencensus.io"
)
```

#### Example #3 `sigs.k8s.io/yaml@v1.3.0`:
```golang
// Currently using sigs.k8s.io/yaml v1.3.0 
const (
    Yaml_K8s_Version = "v1.3.0"
    Yaml_K8s_Path    = "sigs.k8s.io/yaml"
)
```
</details>


#### Example #1: Bash piping
```shell
go list -m -json all | genver 
```

#### Example #2: Read entire contents as string
```shell
genver "$(go list -m -json all)"
```

#### Flags 
All are optional and have reasonable defaults:
* `--out`: directs the location of the generated code. (default `"versions.gen.go"``)
* `--package`: names the package of the generated code. (default `"genver"``)
* `--validate`: validate the generated file against the [Go AST parser](https://pkg.go.dev/go/parser#ParseFile) (default `true`)

### Package

The package exposes a `NewDependencyVersionStore` function that returns a `*DependencyStore`. This will pull in dependencies from runtime and populate each module name and version into a [trie](https://en.wikipedia.org/wiki/Trie). This exposes the following methods:
* `FindVersionFromPath(path string) any`: Pass in an arbitrary pathname and get its version as output (or `nil` if not found)
* `FindVersionFromData(data any) any`: Pass in an arbitrary data type and the version of the module that imported it in get the module (or `nil` if not found)

#### TODO:
- Provide more code examples
- Potentially fix installation guide
- Give more rationale behind using a trie
- Explain more on rationale behind creation
- Add makefile and github actions
- Add more tests and documentation -- perhaps consolidate all code into single `internal` package
