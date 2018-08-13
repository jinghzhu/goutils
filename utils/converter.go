package utils

import (
    "encoding/json"
    "fmt"
    "strconv"
)

// ToString converts input to string.
func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
  
	return string(res)
}

// ToJSON converts input to JSON string.
func ToJSON(obj interface{}) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		return []byte(""), err
	}
	
	return string(res), nil
}

// ToFloat converts string to float, or 0.0 if the input is not a float.
func ToFloat(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}
	
	return res, err
}
