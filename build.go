package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

var templ = `package {{.PkgName}}

import (
{{ range .Imports }}
    {{ if .Name }}
    {{.Name}} {{.Path}}
    {{ else }}
    {{.Path}}
    {{ end }}
{{ end }}
)

{{ range .Constructors }}
// {{.Name}} Create a new {{.Struct}}
func {{.Name}}({{.Params}}) *{{.Struct}} {
    return &{{.Struct}} {
        {{.Fields}}
    }
}
{{ end }}
`

func generateCode(pkgName string, importResuts []ResultImport, results []Result) (string, error) {
	t, err := template.New("").Parse(templ)
	if err != nil {
		return "", err
	}
	data := o{
		"PkgName":     pkgName,
		"Imports":     importResuts,
		"Constructor": []o{},
	}
	constructors := []o{}
	for _, result := range results {
		params := []string{}
		fields := []string{}
		for _, field := range result.Fields {
			params = append(params, fmt.Sprintf("%v %v", field.Name, field.Type))
			fields = append(fields, fmt.Sprintf("%v: %v,", field.Name, field.Name))
		}
		constructors = append(constructors, o{
			"Name":   "New" + result.StructName,
			"Struct": result.StructName,
			"Params": strings.Join(params, ", "),
			"Fields": strings.Join(fields, "\n"),
		})
	}
	data["Constructors"] = constructors

	var buffer bytes.Buffer
	err = t.Execute(&buffer, data)
	if err != nil {
		return "", err
	}
	buf, err := FormatSource(buffer.Bytes())
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func FormatSource(source []byte) ([]byte, error) {
	return imports.Process("", source, &imports.Options{
		AllErrors:  true,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
		FormatOnly: true,
	})
}

type o map[string]interface{}
