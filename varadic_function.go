package main

import "fmt"

func main() {
	fmt.Println(sumAll(10, 10, 2, 1, 3, 4, 10, 2, 11, 1, 1, 1, 123, 32, 32, 3))
	numbers := []int{10, 1, 2, 2, 42, 2, 2, 34, 2}
	fmt.Println(sumAll(numbers...))
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}
