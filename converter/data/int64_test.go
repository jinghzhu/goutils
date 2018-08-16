package data

import (
	"testing"
	"math"
)

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

func TestStrToInt64(t *testing.T) {
	validInts := []int64{0, 1, -1, math.MaxInt8, math.MinInt8, math.MaxInt16, math.MinInt16, math.MinInt32, math.MaxInt32, math.MaxInt64, math.MinInt64}
	invalidInts := []string{"1.233", "a", "false"}

	for _, f := range validInts {
		c, err := StrToInt64(FormatInt64(f))
		if assert.NoError(t, err) {
			assert.EqualValues(t, f, c)
		}
	}
	for _, f := range invalidInts {
		_, err := StrToInt64(f)
		assert.Error(t, err, "expected '"+f+"' to generate an error")
	}
}