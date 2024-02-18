package symbol

import "strconv"

/*
DecToBin converts a decimal to binary
*/
func DecToBin(value int) string {
	return strconv.FormatInt(int64(value), 2)
}
