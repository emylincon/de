package environment

import (
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

func TestDelete(t *testing.T) {
	err := os.Chdir(directory)
	if err != nil {
		log.Fatalf("TestDelete: [ChDir Error]:%v \n", err)
	}
	err = pkgm.Delete()
	if err != nil {
		t.Errorf("TestDelete: [Error]: %v", err)
	}
	err = os.Chdir("../")
	if err != nil {
		log.Fatalf("TestDelete: [ChDir Error]:%v \n", err)
	}
}
