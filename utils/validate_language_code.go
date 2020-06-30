package utils

import "regexp"

// ValidateLanguageCode - Validates a language code.
func ValidateLanguageCode(code string) (valid bool) {
	rx := regexp.MustCompile(`^[a-z]{2,2}-[A-Z]{2,2}$`)

	valid = rx.MatchString(code)
	return
}
