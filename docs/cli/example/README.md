## Example

Example project showing [sirupsen/logrus](https://github.com/sirupsen/logrus) logging its own version -- with a GenVer file being used to retrieve info.

### Usage
This example requires the `genver` command-line tool as well as `make` be installed. 

Run `make all` to automatically generate the version file, compile the source code, and run the service -- which should be followed by:

```shell
INFO[0000] Currently using logrus github.com/sirupsen/logrus@v1.9.3 
```

#### GenVer CLI
Generated using the `genver` cli with the `--package` flag set to `main` from `make gen`. See the output in [versions.gen.go](./versions.gen.go).