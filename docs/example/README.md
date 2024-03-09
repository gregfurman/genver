## Example

Example project showing [sirupsen/logrus](https://github.com/sirupsen/logrus) logging its own version -- with a GenVer file being used to retrieve info.

### Usage
This example requires the `genver` command-line tool as well as `make` be installed. Use `go generate` or `make gen` to create/update the `versions.gen.go` file.

Once generated, proceed with a `go run` (including the generated file) in compilation:

```shell
go run main.go versions.gen.go
```

Alternatively `make all` will automatically generate and compile the source code -- which should be followed by:

```shell
INFO[0000] Currently using logrus github.com/sirupsen/logrus@v1.9.3 
```


#### GenVer CLI
Generated using the `genver` cli with the `--package` flag set to `main`. See the output in [versions.gen.go](./versions-gen.go).

Alternatively, run `make gen`, `go generate`, or the below:
```shell
go list -m -json all | genver --package main
```
