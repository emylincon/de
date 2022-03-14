package environment

import (
	"errors"
	"log"
	"os"
	"testing"
)

var (
	pkgm      = NewPythonVenvMgr()
	directory = "emeka"
)

func setUp() {
	err := pkgm.Create(directory)
	if err != nil {
		log.Fatalf("SetUp Error:%v \n", err)
	}
}

func tearDown() {
	err := os.RemoveAll(directory)
	if err != nil {
		log.Fatalf("tearDown Error:%v \n", err)
	}

	pkgm.Delete()
}

func before() {
	err := os.Chdir(directory)
	if err != nil {
		log.Fatalf("[ChDir Error]: %v", err)
	}
}

func after() {
	err := os.Chdir("../")
	if err != nil {
		log.Fatalf("[ChDir Error]: %v", err)
	}
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestPythonVenv(t *testing.T) {
	path := directory + "/venv"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Error TestPythonVenv: \n[Path]: %v \n[Err]: %v", path, err)
	}
}

func TestRequirements(t *testing.T) {
	path := directory + "/requirements.txt"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Error TestRequirements: \n[Path]: %v \n[Err]: %v", path, err)
	}

}

func TestActivatePythonEnv(t *testing.T) {
	before()
	if _, err := pkgm.Activate(); err != nil {
		t.Errorf("TestActivatePythonEnv: %v", err)
	}
	after()
}

func TestDeactivatePythonEnv(t *testing.T) {
	before()
	if response := pkgm.Deactivate(); response != "Run command on your shell:\n\033[33mdeactivate \033[0m" {
		t.Errorf("TestDeactivatePythonEnv: Response do not match: %v", response)
	}
	after()
}

func TestDelete(t *testing.T) {
	before()
	err := pkgm.Delete()
	if err != nil {
		t.Errorf("TestDelete: [Error]: %v", err)
	}
	after()
}
