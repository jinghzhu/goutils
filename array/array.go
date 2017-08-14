package array

func Index(arr []string, s string) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}

func Include(arr []string, s string) bool {
	return Index(arr, s) >= 0
}
