package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func main() {
	text := "Rizal"
	HelloWithFilter(text, filterText)
}

func HelloWithFilter(text string, filter Filter) {
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
