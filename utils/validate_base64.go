package utils

import "regexp"

// ValidateBase64 checks if a string is a valid base64.
func ValidateBase64(s string) (valid bool) {
	re := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)

	if re.MatchString(s) {
		valid = true
	}
	return
}
