package cmd

import (
	"errors"
	"os"
	"text/template"

	"github.com/emylincon/de/cmd/templates"
)

// Details struct
type Details struct {
	Name  string
	Email string
}

func createEnvironment(path, name, email, kind string) error {
	tmp := ""

	switch kind {
	case "python":
		path += "/main.py"
		tmp = templates.PythonTemplate()
	case "golang":
		path += "/main.go"
		tmp = templates.GolangTemplate()
	default:
		return errors.New("kind can either be 'golang' or 'python'")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	d := Details{
		Name:  name,
		Email: email,
	}

	tmpl, err := template.New("temp").Parse(tmp)
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, d)
	return err
}
