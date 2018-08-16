package data

import "testing"

func TestStrToFloat(t *testing.T) {
	testCases := []string{"", "1", "-.01342", "10.", "empty", "1.23e3", ".23e10"}
	expected := []float64{0, 1, -0.01342, 10.0, 0, 1230, 0.23e10}
	for k, v := range testCases {
		res, _ := StrToFloat(v)
		if res != expected[k] {
			t.Log("Case ", k, ": expected ", expected[k], " when result is ", res)
			t.FailNow()
		}
	}
}
