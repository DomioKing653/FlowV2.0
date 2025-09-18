package shared

import (
    "fmt"
    "os"
)

func Check(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	} else {
		return
	}
}
