package functions

import (
	"fmt"
)

func Check(err error) {
	if err != nil {
		fmt.Print(err)
		if Console {
			mainProgram(publicCode)
		}
	} else {
		return
	}
}
