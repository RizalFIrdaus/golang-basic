package main

import "fmt"

func main() {
	var name = "Hello Brok"
	fmt.Println(name)

	name = "Guooblogg"
	fmt.Println(name)

	datePlace := "Banyuwangi"
	fmt.Println(datePlace)

	//Multiple Variable
	var (
		firstnName = "Muhammaad"
		middleName = "Rizal"
		lastName   = "Firdaus"
	)

	fmt.Println(firstnName, middleName, lastName)

	//	Constant

	const companyName = "Master Bagasi"

	fmt.Println(companyName)
}
