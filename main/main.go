package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinghzhu/goutils/utils"
)

var data int = 1

func main() {
	done, err := utils.Retry(3*time.Second, 3, func() (bool, error) {
		if data == 2 {
			return false, errors.New("error")
		}
		data++

		return false, nil
	})
	if done {
		fmt.Println("done")
	}
	if err != nil {
		fmt.Println("error")
	}
}
