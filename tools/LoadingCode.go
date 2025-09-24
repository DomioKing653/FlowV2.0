package tools

import (
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(filename string) string {
	PublicPath = filename
	code, err := os.ReadFile(filename)
	checkErr(err)
	return string(code)
}
