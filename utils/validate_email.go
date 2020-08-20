package utils

import "regexp"

// ValidateEmail validates a email address.
func ValidateEmail(email string) (valid bool) {
	if email == "" {
		return
	}

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	valid = re.MatchString(email)
	return
}
