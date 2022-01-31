package main

import (
	"flag"
	"log"
	"os"
)

// TODO
func validateName(name *string) (vname *string, err error) {
	return name, err
}

// TODO
func handlePython(cmd *flag.FlagSet, name *string) {
	cmd.Parse(os.Args[2:])
	log.Println("python:", *name)
}

// TODO
func handleGolang(cmd *flag.FlagSet, name *string) {
	cmd.Parse(os.Args[2:])
	log.Println("golang:", *name)
}

func main() {
	usage := "create-python-app or create-go-app"
	python := flag.NewFlagSet("create-python-app", flag.ExitOnError)
	golang := flag.NewFlagSet("create-go-app", flag.ExitOnError)

	papp := python.String("app", "myapp", "app name")
	gapp := golang.String("app", "myapp", "app name")

	if len(os.Args) < 1 {
		log.Fatalln(usage)
	}
	gapp, err := validateName(gapp)
	if err != nil {
		log.Fatalln(err)
	}
	papp, err = validateName(papp)
	if err != nil {
		log.Fatalln(err)
	}
	switch os.Args[1] {
	case "create-python-app":
		handlePython(python, papp)
	case "create-go-app":
		handleGolang(golang, gapp)
	default:
		log.Fatalln(usage)
	}

}
