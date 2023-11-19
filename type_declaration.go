package main

import "fmt"

func main() {

	type npm string

	var rizal npm = "54418907"

	var mifmelse string = "544111324"
	var cNpm npm = npm(mifmelse)

	fmt.Println(rizal)
	fmt.Println(cNpm)
}
