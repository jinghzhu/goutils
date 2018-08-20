package utils

import (
	"encoding/base64"
	"io/ioutil"
)

// EncodeBase64 accepts a file path. Return its base64 code and any error.
func EncodeBase64(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buff), nil
}

// DecodeBase64 accepts the string of base64 code and a file path. It decodes
// the string to that file. Return error if any issue.
func DecodeBase64(code string, dest string) error {
	buff, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(dest, buff, 07440); err != nil {
		return err
	}

	return nil
}
