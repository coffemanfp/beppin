package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetFilebytes - Gets the bytes from a file by its filepath.
func GetFilebytes(path string) (fileBytes []byte, err error) {
	if path == "" {
		return
	}

	exists, err := ExistsFile(path)
	if err != nil {
		return
	}

	if !exists {
		err = fmt.Errorf("%w: %s", os.ErrNotExist, path)
		return
	}

	fileBytes, err = ioutil.ReadFile(path)
	return
}
