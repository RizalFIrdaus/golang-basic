package main

import "fmt"

func random()any{
	return true
}

func main() {
	var result any = random()

	switch value := result.(type) {
	case int:
		fmt.Println("ini integer mas :", value)
	case string:
		fmt.Println("ini string mas :", value)
	default: 
	fmt.Println("Salah boy")
	}
}