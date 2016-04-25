package generator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/guillaumebreton/regen/loader"
)

// Generator from a template
type Generator struct {
	template *template.Template
}

// NewGenerator Creates a new generator
func NewGenerator(templatePath string) (*Generator, error) {
	t, err := template.ParseFiles(templatePath)

	if err != nil {
		return nil, fmt.Errorf("Fail to load template %s, %s", templatePath, err)
	}
	return &Generator{t}, nil
}

// Execute generates a resume using data and a template path
func (g *Generator) Execute(resume *loader.Resume) (string, error) {

	buf := new(bytes.Buffer)
	err := g.template.Execute(buf, resume)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
