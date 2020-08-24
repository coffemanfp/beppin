package utils

import "strconv"

// ParseUint - Converts a string to a int.
func ParseUint(s string, bitSize int) (i uint64, err error) {
	if s == "" || s == " " {
		return
	}

	i, err = strconv.ParseUint(s, 10, bitSize)
	return
}
