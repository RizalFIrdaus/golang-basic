package main

import (
	formatter_response "basic/helper"
	"fmt"
)

func main(){
	result := formatter_response.Success("200", "Success")
	fmt.Println(result)

	test := formatter_response.GetVersion()
	formatter_response.SetVersion("2")
	test2 := formatter_response.GetVersion()
	fmt.Println(test)
	fmt.Println(test2)

}