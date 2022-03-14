package environment

import (
	"os"
	"os/exec"
)

// CreateGoEnvironment runs go mod tidy and create .gitignore
func CreateGoEnvironment(directory string) error {
	if err := goModTidy(directory); err != nil {
		return err
	}
	return getIgnore(directory+"/.gitignore", "golang")

}

// goModTidy runs go mod tidy to init the module
func goModTidy(directory string) error {
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
