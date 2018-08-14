package utils

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
