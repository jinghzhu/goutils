package data

// Uint returns a pouinter to of the uint value passed in.
func Uint(v uint) *uint {
	return &v
}

// UintVal returns the value of the uint pouinter passed in or
// 0 if the pouinter is nil.
func UintVal(v *uint) uint {
	if v != nil {
		return *v
	}

	return 0
}

// UintSlice converts a slice of uint values uinto a slice of
// uint pouinters
func UintSlice(src []uint) []*uint {
	dst := make([]*uint, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}

	return dst
}

// Int64 returns a pointer to of int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Val returns the value of int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Val(v *int64) int64 {
	if v != nil {
		return *v
	}

	return 0
}

// Int64Slice converts a slice of int64 values into a slice of
// int64 pointers
func Int64Slice(src []int64) []*int64 {
	dst := make([]*int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}

	return dst
}

// Int64ValSlice converts a slice of int64 pointers into a slice of
// int64 values
func Int64ValSlice(src []*int64) []int64 {
	dst := make([]int64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}

	return dst
}

// Int64Map converts a string map of int64 values into a string
// map of int64 pointers
func Int64Map(src map[string]int64) map[string]*int64 {
	dst := make(map[string]*int64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}

	return dst
}

// Int64ValMap converts a string map of int64 pointers into a string
// map of int64 values
func Int64ValMap(src map[string]*int64) map[string]int64 {
	dst := make(map[string]int64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}

	return dst
}
