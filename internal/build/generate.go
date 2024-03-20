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

	buildInfo := GetBuildInfoFromData(data)
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, struct {
		Package       string
		GenVerVersion string
		Dependencies  []DepInfo
	}{
		Package:            pkg,
		Dependencies:       buildInfo.Dependencies,
		VersionInformation: buildInfo.VersionInfo,
	}); err != nil {
		return err
	}

	if mustValidate {
		if err := isValidGo(tpl.String()); err != nil {
			return err
		}
	}

	if err := os.WriteFile(outPath, tpl.Bytes(), 0644); err != nil {
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
