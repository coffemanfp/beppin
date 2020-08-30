package models

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"path"
	"strings"

	"github.com/coffemanfp/beppin-server/config"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// Avatar represents a user avatar
type Avatar struct {
	URL  string `json:"url,omitempty"`
	Data string `json:"data,omitempty"`
}

// Save saves the avatar on the file system
func (a Avatar) Save(userIdentifier string) (avatarURL string, err error) {
	if !a.ValidateData() {
		err = fmt.Errorf("failed to validate avatar: %w (avatar)", errs.ErrInvalidData)
		return
	}

	avatarPath := path.Join(
		config.GetVar("assets").(string),
		"avatars",
		userIdentifier,
	)

	err = ioutil.WriteFile(avatarPath, []byte(a.Data), 0777)
	if err != nil {
		err = fmt.Errorf("failed to write avatar file: %v", err)
	}

	avatarURLToParse, err := url.Parse(config.GetVar("host").(string))
	if err != nil {
		err = fmt.Errorf("failed to parse host url: %v", err)
		return
	}

	avatarURLToParse.Path = path.Join(avatarURLToParse.Path, avatarPath)
	avatarURL = avatarURLToParse.String()
	return
}

// Validate Validates a avatar
func (a Avatar) Validate() (valid bool) {
	valid = true

	if a.URL == "" && a.Data == "" {
		valid = false
	}

	if a.URL != "" && !a.ValidateURL() {
		valid = false
		return
	}

	if a.Data != "" && !a.ValidateData() {
		valid = false
	}
	return
}

// ValidateURL checks if a string is a valid url
func (a Avatar) ValidateURL() (valid bool) {
	valid = true

	if a.URL == "" {
		valid = false
		return
	}

	if _, err := url.ParseRequestURI(a.URL); err != nil {
		valid = false
	}
	return
}

// ValidateData checks if a string is a valid avatar data file
func (a Avatar) ValidateData() (valid bool) {
	valid = true

	if a.Data == "" {
		valid = false
		return
	}

	var justBase64 string

	contentType := a.Data[:30]
	if strings.Contains(contentType, "image/png") {
		justBase64 = a.Data[22:]
	} else if strings.Contains(contentType, "image/jpeg") {
		justBase64 = a.Data[23:]
	} else {
		valid = false
		return
	}

	_, err := base64.StdEncoding.DecodeString(justBase64)
	if err != nil {
		valid = false
	}

	return
}
