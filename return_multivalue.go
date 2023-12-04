package main

import "fmt"

func main() {
	first, _ := sum(2, 2)

	fmt.Println(first)
}

func sum(a int, b int) (int, string) {
	return a + b, "Hello Brok"
}
