package data

import "strconv"

// StrToFloat converts string to float, or 0.0 if the input is not a float.
func StrToFloat(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}

	return res, err
}
