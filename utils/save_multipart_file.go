package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	errs "github.com/coffemanfp/beppin/errors"
)

// SaveMultipartFile - Saves a multipart file.
func SaveMultipartFile(file *multipart.FileHeader, destination string, maxSize int) (err error) {
	// If the maxSize is provided, check if the file size is greater that
	// the maxSize
	if maxSize > 0 && int(file.Size) > maxSize {
		err = fmt.Errorf("%w: size exceeded", errs.ErrInvalidData)
		return
	}

	src, err := file.Open()
	if err != nil {
		err = fmt.Errorf("%w: corrupt file", errs.ErrInvalidData)
		return
	}
	defer src.Close()

	var dest *os.File

	if exists, _ := ExistsFile(destination); exists {
		dest, err = os.OpenFile(destination, os.O_WRONLY, 0777)
	} else {
		dest, err = os.Create(destination)
	}
	if err != nil {
		return
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return
}
