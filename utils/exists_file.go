package utils

import (
	"fmt"
	"os"
)

// ExistsFile - Checks if exists a file.
func ExistsFile(path string) (exists bool, err error) {
	if path == "" {
		return
	}

	exists = true

	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = nil
			exists = false
			return
		} else {
			err = fmt.Errorf("failed to check (%s) file: %v", path, err)
			exists = false
			return
		}
	}

	return
}
