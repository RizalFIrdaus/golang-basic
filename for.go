package main

import "fmt"

func main() {

	//for counter := 1; counter <= 100; counter++ {
	//	fmt.Println("Perulangan ke ", counter)
	//}
	//
	//fmt.Println("Selesai")f

	// For Range
	names := []string{"Muhammad", "Rizal", "Firdaus"}

	for i := 0; i < len(names); i++ {
		fmt.Println(i+1, names[i])
	}
	fmt.Println("")
	for i, name := range names {
		fmt.Println(i+1, name)
	}
	fmt.Println("")
	for _, name := range names {
		fmt.Println(name)
	}

}
