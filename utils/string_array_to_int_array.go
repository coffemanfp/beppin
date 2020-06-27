package utils

import (
	"fmt"
	"strconv"
)

// StringArrayToIntArray - Convert a string array to an int array.
//	@param sArray []string:
//		Array to convert.
//	@return iArray []int:
//		Converted Array.
func StringArrayToIntArray(sArray []string) (iArray []int, err error) {
	var i int

	for _, s := range sArray {
		i, err = strconv.Atoi(s)
		if err != nil {
			err = fmt.Errorf("failed to convert %s string to an int:\n%s", s, err)
			return
		}

		iArray = append(iArray, i)
	}
	return
}
