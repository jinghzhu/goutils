package data

import "strconv"

// UintPointer returns a pointer to of the uint value passed in.
func UintPointer(v uint) *uint {
	return &v
}

// UintVal returns the value of the uint pointer passed in or
// 0 if the pointer is nil.
func UintVal(v *uint) uint {
	if v != nil {
		return *v
	}

	return 0
}

// UintSlice converts a slice of uint values uinto a slice of
// uint pointers
func UintSlice(src []uint) []*uint {
	dst := make([]*uint, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}

	return dst
}

// UintValSlice converts a slice of uint pointers uinto a slice of
// uint values
func UintValSlice(src []*uint) []uint {
	dst := make([]uint, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// StrToUint64 turn a string into a uint64.
func StrToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
