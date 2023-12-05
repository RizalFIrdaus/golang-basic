package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "ANJING"
	HelloWithFilter(text, filterText)
}

func HelloWithFilter(text string, filter func(string) string) {
	filterText := filter(text)
	fmt.Println(filterText)
}

func filterText(text string) string {
	lowerCase := strings.ToLower(text)
	if lowerCase == "anjing" {
		return "..."
	}
	return lowerCase
}
