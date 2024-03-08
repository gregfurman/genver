package build

import (
	"embed"
	"text/template"
)

//go:embed template.tmpl
var tmplFile embed.FS

func retrieveTemplate() (*template.Template, error) {
	templateContent, err := tmplFile.ReadFile("template.tmpl")
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("template.tmpl").Parse(string(templateContent))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
