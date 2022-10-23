//go:generate go run generr.go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v3"
)

type ErrorDef struct {
	Name    string
	Message string
	Vars    []ErrorVar
}

type ErrorVar struct {
	Name string
	Type string
}

var tmp = `
// generate code - DO NOT EDIT

package errs


type UserError struct {
	Name string
	Msg string
}

func (e *UserError) Error () string {
	return e.Msg
}

{{ range . }}
func {{ .Name }}({{ range .Vars }} {{ .Name }} {{ .Type }}, {{ end }}) error {
	return &UserError{
		Name: "{{ .Name }}",
		Msg: fmt.Sprintf("{{ .Message }}", {{ range .Vars }} {{ .Name }}, {{ end }}),
	}
}
{{ end }}

`

func main() {
	var ed []ErrorDef
	b, err := ioutil.ReadFile("../errs/errdef.yaml")
	if err != nil {
		log.Fatalf("read file err: %v", err)
	}

	err = yaml.Unmarshal(b, &ed)
	if err != nil {
		log.Fatalf("unmarshal err: %v\n", err)
	}

	t, err := template.New("errgens").Parse(tmp)
	if err != nil {
		log.Fatalf("parse tempalte err: %v", err)
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, ed)
	if err != nil {
		log.Fatalf("excute template err: %v", err)
	}

	code, err := imports.Process("", buf.Bytes(), &imports.Options{
		Comments: true,
	})

	if err != nil {
		log.Fatalf("format code err: %v", err)
	}

	err = os.WriteFile("../errs/err.gen.go", code, os.ModePerm)
	if err != nil {
		log.Fatalf("write file err: %v", err)
	}

	fmt.Println("generate file success")
}
