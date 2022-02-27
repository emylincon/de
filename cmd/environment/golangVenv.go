package environment

import (
	"os"
	"os/exec"
)

// GoModTidy runs go mod tidy to init the module
func GoModTidy(directory string) error {
	workingDir, _ := os.Getwd()
	if err := os.Chdir(directory); err != nil {
		return err
	}
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		return err
	}
	if err := os.Chdir(workingDir); err != nil {
		return err
	}
	return nil
}
