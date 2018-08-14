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
