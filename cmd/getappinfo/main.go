package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/exelestor/godev_test_task/pkg/appinfo"
)

var lang = ""

type Output struct {
	Result []appinfo.Result `json:"result,omitempty"`
}

func (o *Output) AppendLoop(in chan appinfo.Result) {
	for {
		o.Result = append(o.Result, <-in)
	}
}

func (o *Output) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

func scanFile(f *os.File, apps chan string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		apps <- scanner.Text()
	}
	close(apps)
}

func main() {
	processArguments()

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	var workers Workers
	workers.Init()
	workers.Run()

	var output Output
	go output.AppendLoop(workers.ProcessedApps)

	scanFile(file, workers.AppsToProcess)

	workers.WaitUntilDone()

	j, _ := json.Marshal(output)
	fmt.Println(string(j))

	//fmt.Println(output)
}
