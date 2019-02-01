package main

import (
	"encoding/json"
	"fmt"
	"github.com/exelestor/godev_test_task/pkg/appinfo"
)

func main() {
	res, err := appinfo.Get("air.com.toshi.ppkp", "")
	if err != nil {
		fmt.Println(err)
	}

	out, _ := json.Marshal(res)
	fmt.Println(string(out))
}
