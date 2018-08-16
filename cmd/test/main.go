package main

import (
	"fmt"
	"time"

	"github.com/jinghzhu/goutils/http"
)

func main() {
	done := make(chan bool, 1)
	go doneFunc(done)

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Reach timeout")
	case <-done:
		fmt.Println("Reach done")
	}

	done <- true
	time.Sleep(15 * time.Second)
}

func doneFunc(ch chan bool) {
	defer close(ch)
	resp, err := http.HttpGet("https://ldap.svc.eng.vmware.com/v1.0/user/jianlanz", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
	respStr, err := http.ResponseToString(resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(respStr)
	time.Sleep(5 * time.Second)
}
