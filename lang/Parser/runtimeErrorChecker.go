package Parser

import "fmt"

func CheckRuntimeErr(err error){
	if err!=nil{
		fmt.Println("Runtime error:"+err.Error())
	}else {
		return
	}
}