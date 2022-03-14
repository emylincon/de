package environment

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func getIgnore(filepath, kind string) error {
	var url string
	if kind == "python" {
		url = "https://github.com/github/gitignore/raw/main/Python.gitignore"
	} else if kind == "golang" {
		url = "https://github.com/github/gitignore/raw/main/Go.gitignore"
	} else {
		return errors.New("unknown kind")
	}

	return downloadFile(filepath, url)
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
