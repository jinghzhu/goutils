package converter

import "strconv"

// StrToInt64 turns an int64 into a string
func StrToInt64(value int64) string {
	return strconv.FormatInt(value, 10)
}
