package utils

import (
	"bytes"
	"net/http"
)

// GetFileContentType - Gets the file content type.
func GetFileContentType(out bytes.Buffer) (contentType string, err error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType = http.DetectContentType(buffer)
	return
}
