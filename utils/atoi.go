package utils

import "strconv"

// Atoi - Converts a string to a int.
func Atoi(s string) (i int, err error) {
	if s != "" || s == " " {
		i, err = strconv.Atoi(s)
		if err != nil {
			return
		}
	}

	return
}
