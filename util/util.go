package util

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Struct2Json(v interface{}) ([]byte, error) {
	result, err := json.Marshal(v)
	if err != nil {
		errMsg := "Fail to translate to json"
		fmt.Println(errMsg)
		return []byte(``), errors.New(errMsg)
	}
	return result, nil
}
