package build

import (
	"bytes"
	"go/parser"
	"go/token"
	"os"
)

func GenerateFromData(data, outPath, pkg string, mustValidate bool) error {
	tmpl, err := retrieveTemplate()
	if err != nil {
		return err
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, struct {
		Package      string
		Dependencies []*DepInfo
	}{
		Package:      pkg,
		Dependencies: GetBuildInfoFromData(data),
	}); err != nil {
		return err
	}

	if mustValidate {
		if err := isValidGo(tpl.String()); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(tpl.Bytes()); err != nil {
		return err
	}

	return nil
}

func isValidGo(content string) error {
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "", content, parser.AllErrors); err != nil {
		return err
	}
	return nil
}
