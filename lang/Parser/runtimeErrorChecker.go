package Parser

import (
	"fmt"
	"os"
)

func CheckRuntimeErr(err error) {
	if err != nil {
		fmt.Println("Runtime error!")
		fmt.Println("┬───────────────────────────────────────────")
		fmt.Println("│")
		fmt.Println("│" + err.Error())
		fmt.Println("│")
		fmt.Println("┴───────────────────────────────────────────")
		os.Exit(-2)
	} else {
		return
	}
}
