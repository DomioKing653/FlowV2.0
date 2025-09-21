package shared

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Print("Synatax error:" + err.Error())
		os.Exit(-1)
	} else {
		return
	}
}
