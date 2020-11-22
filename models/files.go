package models

import (
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/coffemanfp/beppin/config"
)

// File - File representation.
type File struct {
	ID int64 `json:"id,omitempty"`

	Path string `json:"path,omitempty"`
	URL  string `json:"url,omitempty"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Files - Alias for a file array.
type Files []File

// SetURL join the host with the f.Path value and return
// the result
func (f *File) SetURL() (err error) {
	u, err := url.Parse(config.GlobalSettings.Host)
	if err != nil {
		return
	}

	u.Path = path.Join(
		u.Path,
		f.Path,
	)
	f.URL = u.String()
	return
}

// GetIdentifier gets the first unique identifier it finds in order of importance.
func (f File) GetIdentifier() (identifier interface{}) {
	if f.ID != 0 {
		identifier = f.ID
	} else if f.Path != "" {
		identifier = f.Path
	}

	return
}

// NewFilePath - returns a new file path.
func NewFilePath(name string) (newPath string) {
	prefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	newPath = path.Join(
		config.GlobalSettings.Assets,
		prefix+name,
	)
	return
}
