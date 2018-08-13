package utils

import (
    "fmt"
)

// ToString convert the input to a string.
func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
  
	return string(res)
}
