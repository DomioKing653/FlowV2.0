package Parser

import (
	"fmt"
	"os"
)

func CheckRuntimeErr(err error) {
	if err != nil {
		fmt.Println("Runtime error:" + err.Error())
		os.Exit(-2)
	} else {
		return
	}
}
