package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"log/slog"

	"github.com/gregfurman/genver/internal/build"
)

func main() {
	fout := flag.String("out", "versions.gen.go", "location of generated dependency information.")
	fpackage := flag.String("package", "genver", "package of generated dependency information.")
	fvalidateProtoSyntax := flag.Bool("validate", true, "if enabled, uses the Go AST parser to check if the generated file has valid syntax")
	flag.Parse()

	args := flag.Args()

	content, err := processArgs(args)
	if err != nil {
		slog.Error("failed to read data from stdin", slog.Any("error", err))
		return
	}

	if err := run(content, *fout, *fpackage, *fvalidateProtoSyntax); err != nil {
		slog.Error("failed to create genver file", slog.Any("err", err))
		return
	}

	slog.Info("successfully created genver file", slog.Any("location", *fout))
}

func processArgs(args []string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, ""), nil
	}

	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, os.Stdin)
	if err != nil {
		return "", err
	} else if n <= 1 {
		return "", fmt.Errorf("no input found")
	}

	return buf.String(), nil
}

func run(data, out, pkg string, validateSyntax bool) error {
	if err := build.GenerateFromData(data, out, pkg, validateSyntax); err != nil {
		return err
	}
	return nil
}
