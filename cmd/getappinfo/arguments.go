package main

import (
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Printf("Usage: %s FILE [LANG]\n", os.Args[0])
	os.Exit(0)
}

func processArguments() {
	if len(os.Args) != 2 {
		if len(os.Args) == 3 {
			lang = strings.ToLower(os.Args[2])
		} else {
			printUsage()
		}
	}
}
