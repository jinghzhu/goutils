package main

import (
	"fmt"

	"github.com/jinghzhu/goutils/utils"
)

func main() {
	str, err := utils.EncodeBase64("/Users/jinghuaz/toolchain/k8s/k8s_config/gulel_nimbus_worker_staging/config")
	if err != nil {
		panic(err)
	}

	fmt.Println(str)
}
