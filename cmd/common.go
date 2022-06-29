package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/emylincon/dec/cmd/licences"
	"github.com/emylincon/dec/cmd/templates"
)

// Details struct
type Details struct {
	Name  string
	Email string
}

// LicenseObj struct
type LicenseObj struct {
	Fullname string
	Year     int
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
	tmpData := Details{
		Name:  name,
		Email: email,
	}

	tmpl, err := template.New("temp").Parse(tmp)
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, tmpData)
	return err
}

func setUpLicense(path, licenseName, fullname string) error {
	tmp := ""
	path += "/LICENSE"

	switch strings.ToLower(licenseName) {
	case "mit":
		tmp = licences.MITTemplate()
	case "apache":
		tmp = licences.ApacheTemplate()
	default:
		return fmt.Errorf("license name: '%s' is not implemented", licenseName)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	tmpData := LicenseObj{
		Fullname: fullname,
		Year:     time.Now().Year(),
	}

	tmpl, err := template.New("temp").Parse(tmp)
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, tmpData)
	return err

}
