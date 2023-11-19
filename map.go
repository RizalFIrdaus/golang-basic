package main

import "fmt"

func main() {

	person := map[string]string{
		"name": "Muhammad Rizal Firdaus",
		"age":  "23",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])

	var mapping = make(map[string]string)

	mapping["name"] = "Esan"
	mapping["address"] = "Jakarta"
	mapping["error"] = "error"

	fmt.Println(mapping)

	delete(mapping, "error")

	fmt.Println(mapping)

}
