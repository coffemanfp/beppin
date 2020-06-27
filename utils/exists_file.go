package utils

import (
	"os"
)

// ExistsFile - Checks if exists a file.
func ExistsFile(filepath string) (exists bool, err error) {
	if filepath == "" {
		return
	}

	exists = true

	if _, err = os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			err = nil
			exists = false
			return
		}
	}

	return
}
