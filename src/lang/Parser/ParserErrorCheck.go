package Parser

import (
	"fmt"
)

func Check(err error) {
	if err != nil {
		fmt.Print(err)
	} else {
		return
	}
}
