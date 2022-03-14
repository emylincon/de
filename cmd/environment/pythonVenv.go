package environment

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/atotto/clipboard"
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

func (p PythonVenvMgr) requirements(directory string) error {
	file := []byte("# run pip freeze > requirements.txt")
	return os.WriteFile(directory+"/requirements.txt", file, 0644)

}

func (p PythonVenvMgr) gitignore(directory string) error {
	return getIgnore(directory+"/.gitignore", "python")
}

func (p PythonVenvMgr) updatePip(directory string) error {
	cmd := exec.Command(p.fullVenvPath(directory)+"/bin/python3", "-m", "pip", "install", "--upgrade", "pip")
	return cmd.Run()
}

// Create creates a venv
func (p PythonVenvMgr) Create(directory string) error {
	cmd := exec.Command("python3", "-m", "venv", p.fullVenvPath(directory))
	if err := cmd.Run(); err != nil {
		return err
	}
	if err := p.requirements(directory); err != nil {
		return err
	}
	if err := p.updatePip(directory); err != nil {
		return err
	}
	return p.gitignore(directory)
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
func (p PythonVenvMgr) Activate() (string, error) {
	// verify path exists
	_, err := p.getVenvPath()
	if err != nil {
		return "", err
	}
	path := "venv"

	if runtime.GOOS == "windows" {
		switch p.windowsShell() {
		case "cmd.exe":
			cmd := fmt.Sprintf("%s/Scripts/activate.bat", path)
			clipboard.WriteAll(cmd)
			return fmt.Sprintf("Run on your shell: `%s` ", cmd), nil
		default:
			cmd := fmt.Sprintf("%s/Scripts/Activate.ps1", path)
			return fmt.Sprintf("Run on your shell: `%s` ", cmd), nil
		}
	} else {
		cmd := fmt.Sprintf("source %s/bin/activate", path)
		clipboard.WriteAll(cmd)
		return fmt.Sprintf("Run command on your shell:\n\033[35m%s \033[0m", cmd), nil
	}

}

// Deactivate deactivates venv
func (p PythonVenvMgr) Deactivate() string {
	cmd := "deactivate"
	clipboard.WriteAll(cmd)
	return fmt.Sprintf("Run command on your shell:\n\033[33m%s \033[0m", cmd)
}
