package environment

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

// error variables block
var (
	ErrVenvDoNotExist = errors.New("venv do not exist")
)

// PythonVenvMgr manages creating, deleting, activating and deactiving
// virtual environments
type PythonVenvMgr struct {
	venv string
}

// NewPythonVenvMgr is a PythonVenvMgr constructor
func NewPythonVenvMgr() PythonVenvMgr {
	return PythonVenvMgr{
		venv: "venv",
	}
}

func (p PythonVenvMgr) fullVenvPath(path string) string {
	return path + "/" + p.venv
}

// Create creates a venv
func (p PythonVenvMgr) Create(directory string) error {
	cmd := exec.Command("python3", "-m", "venv", p.fullVenvPath(directory))
	return cmd.Run()
}

func (p PythonVenvMgr) getVenvPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	venvPath := p.fullVenvPath(path)
	if _, err := os.Stat(venvPath); os.IsNotExist(err) {
		return "", ErrVenvDoNotExist
	}
	return venvPath, nil
}

// Delete deletes venv
func (p PythonVenvMgr) Delete() error {
	path, err := p.getVenvPath()
	if err == ErrVenvDoNotExist {
		return nil
	} else if err != nil {
		return err
	}
	return os.RemoveAll(path)
}

// adapted from https://github.com/riywo/loginshell/blob/master/loginshell.go
func (p PythonVenvMgr) windowsShell() string {
	consoleApp := os.Getenv("COMSPEC")
	if consoleApp == "" {
		consoleApp = "cmd.exe"
	}

	return consoleApp
}

// Activate activates venv
func (p PythonVenvMgr) Activate() error {
	path, err := p.getVenvPath()
	if err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		switch p.windowsShell() {
		case "cmd.exe":
			return exec.Command(path + "/Scripts/activate.bat").Run()
		default:
			return exec.Command(path + "/Scripts/Activate.ps1").Run()
		}
	} else {
		return exec.Command("source", path+"/bin/activate").Run()
	}

}

// Deactivate deactivates venv
func (p PythonVenvMgr) Deactivate() error {
	return exec.Command("deactivate").Run()
}
