package data

import (
	"fmt"
	"reflect"
	"strconv"
)

// InterfaceToInt64 converts input to an integer type 64, or 0 if input is not an integer.
func InterfaceToInt64(value interface{}) (res int64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		res = int64(val.Uint())
	case string:
		if IsInt(val.String()) {
			res, err = strconv.ParseInt(val.String(), 0, 64)
			if err != nil {
				res = 0
			}
		} else {
			err = fmt.Errorf("math: square root of negative number %g", value)
			res = 0
		}
	default:
		err = fmt.Errorf("math: square root of negative number %g", value)
		res = 0
	}

	return
}
