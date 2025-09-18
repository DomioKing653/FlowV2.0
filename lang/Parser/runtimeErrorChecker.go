package Parser

import "fmt"

func CheckRuntimeErr(err error){
	if err!=nil{
		fmt.Println(err)
	}else {
		return
	}
}