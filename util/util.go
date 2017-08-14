package util

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func Struct2String(v interface{}) string {
	result, err := json.Marshal(v)
	if err != nil {
		errMsg := "Fail to translate to json"
		fmt.Println(errMsg)
		return fmt.Sprintf("%v", v)
	}
	return string(result)
}

// GetMountPoints returns all moutpoins in a string array
func GetMountPoints(server string) ([]string, error) {
	b, err := exec.Command("showmount", "-e", server).Output()
	if err != nil {
		fmt.Println("error in showmount: " + err.Error())
		return nil, err
	}
	s := strings.TrimSpace(string(b))
	// The fist line of showmount -e <server> is Exports list on <server>
	firstLine := strings.Index(s, "\n")
	sArr := strings.Split(s[firstLine+1:], "\n")
	for i := 0; i < len(sArr); i++ {
		index := strings.Index(sArr[i], " ")
		temp := sArr[i]
		sArr[i] = temp[:index]
	}
	return sArr, nil
}
