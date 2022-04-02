package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)
	for _, value := range text {
		if strings.Count(text, string(value)) == 1 {
			fmt.Print(string(value))
		}
	}
}
