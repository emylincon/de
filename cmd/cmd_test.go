package cmd

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var testPaths = map[string]string{"python": "ipython", "golang": "igolang"}

func setUp() {
	path, name, email := testPaths["python"], "emeka", "emeka@gmail.com"
	err := python(path, name, email)
	if err != nil {
		log.Fatalln("Python Setup error: " + err.Error())
	}

	path, name, email = testPaths["golang"], "emeka", "emeka@gmail.com"
	err = golang(path, name, email)
	if err != nil {
		log.Fatalln("Golang Setup error: " + err.Error())
	}
	log.Println("set up complete")
}

func tearDown() {
	for _, path := range testPaths {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println("Tear Down Error:", err)
		}
	}

}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestDir(t *testing.T) {
	for _, path := range testPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Error TestDir: \n[Path]: %v \n[Err]: %v", path, err)
		}
	}
}

func TestCheckMainFile(t *testing.T) {
	mains := map[string]string{"python": "main.py", "golang": "main.go"}
	for key, v := range mains {
		path := testPaths[key] + "/" + v
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Error TestCheckMainFile: \n[Path]: %v \n[Err]: %v", path, err)
		}
	}
}

func TestPythonVenv(t *testing.T) {
	path := testPaths["python"] + "/venv"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Error TestPythonVenv: \n[Path]: %v \n[Err]: %v", path, err)
	}

}
