package utils

import "net/url"

// ValidateURL validates a url
func ValidateURL(u string) (valid bool) {
	if _, err := url.ParseRequestURI(u); err == nil {
		valid = true
	}
	return
}
