package generator

import (
	"bytes"
	"text/template"

	"github.com/guillaumebreton/regen/loader"
)

func Execute(templatePath string, resume *loader.Resume) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, resume)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
