package main

import (
	"Flow2.0/functions"
	"os"
)

var silent = false
var programPath string

func main() {
	if len(os.Args) > 1 {
		programPath = os.Args[1]
		if len(os.Args) > 2 {
			if os.Args[2] == "-s" {
				silent = true
			}
		}
	}
	code := functions.ReadFile(programPath)
	functions.Run(code, silent)
}
