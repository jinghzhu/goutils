package data

import "testing"

func TestInterfaceToInt64(t *testing.T) {
	testCases := []string{"1000", "-123", "abcdef", "100000000000000000000000000000000000000000000"}
	expected := []int64{1000, -123, 0, 0}
	for k, v := range testCases {
		result, _ := InterfaceToInt64(v)
		if result != expected[k] {
			t.Log("Case ", k, ": expected ", expected[k], " when result is ", result)
			t.FailNow()
		}
	}
}
