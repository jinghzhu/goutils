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

func TestToJSON(t *testing.T) {
	testCasess := []interface{}{"test", map[string]string{"a": "b", "b": "c"}, func() error { return fmt.Errorf("Error") }}
	expected := [][]string{
		{"\"test\"", "<nil>"},
		{"{\"a\":\"b\",\"b\":\"c\"}", "<nil>"},
		{"", "json: unsupported type: func() error"},
	}
	for i, test := range tests {
		actual, err := ToJSON(test)
		if actual != expected[i][0] {
			t.Errorf("Expected toJSON(%v) to return '%v', got '%v'", test, expected[i][0], actual)
		}
		if fmt.Sprintf("%v", err) != expected[i][1] {
			t.Errorf("Expected error returned from toJSON(%v) to return '%v', got '%v'", test, expected[i][1], fmt.Sprintf("%v", err))
		}
	}
}
