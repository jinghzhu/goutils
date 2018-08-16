package data

import "strconv"

// StrToFloat64 converts string to float, or 0.0 if the input is not a float.
func StrToFloat64(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}

	return res, err
}
