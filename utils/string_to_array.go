package utils

import (
	"errors"
	"regexp"
	"strings"
)

// StringToArray - Convert a string in a array.
// Example: "[1, 2, 3]", "[1,2,3]", "1, 2, 3" or "1,2,3".
//	@param s string:
//		String to convert.
//	@return array []string:
//		Array expected.
func StringToArray(s string) (array []string, err error) {
	re := regexp.MustCompile(`^[0-9,\[\],\,\ ]*$`)
	if !re.MatchString(s) {
		err = errors.New("error: int array invalid")
		return
	}
	if strings.Index(s, "[") == 0 {
		s = s[1 : len(s)-1]
	}
	if strings.Contains(s, " ") {
		s = strings.Replace(s, " ", "", -1)
	}

	array = strings.Split(s, ",")
	return
}
