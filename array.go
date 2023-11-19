package main

import (
	"fmt"
)

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Rizal"
	names[2] = "Firdaus"

	fmt.Println(names[0], names[1], names[2])

	var values = [3]string{
		"kodok",
		"zuma",
		"aja",
	}

	fmt.Println(values[0], values[1], values[2])
	fmt.Println(len(values) + len(names))

	var years = [...]int16{
		2013,
		2014,
		2020,
		2024,
	}

	fmt.Println(years)

}
