package templates

// GolangTemplate returns layout
func GolangTemplate() string {
	layout := `/*
Copyright © 2022 {{ .Name }} <{{ .Email }}>
*/

package main

import "fmt"


func main() {
	fmt.Println("Hello World")
}
	`
	return layout
}
