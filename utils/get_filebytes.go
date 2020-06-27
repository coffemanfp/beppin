package utils

import (
	"fmt"
	"io/ioutil"
)

// GetFilebytes - Gets the bytes from a file by its filepath.
func GetFilebytes(filepath string) (fileBytes []byte, err error) {
	if filepath == "" {
		return
	}

	exists, err := ExistsFile(filepath)
	if err != nil {
		return
	}

	if !exists {
		err = fmt.Errorf("file (%s) not found:\n%s", filepath, err)
	}

	fileBytes, err = ioutil.ReadFile(filepath)
	return
}
