package utils

import (
	"fmt"
	"testing"
)

func TestToFloat(t *testing.T) {
	testCasess := []string{"", "1", "-.01342", "10.", "empty", "1.23e3", ".23e10"}
	expected := []float64{0, 1, -0.01342, 10.0, 0, 1230, 0.23e10}
	for k, v := range testCases {
		res, _ := ToFloat(v)
		if res != expected[k] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}
