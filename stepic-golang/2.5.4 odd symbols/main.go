package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	for i, value := range []rune(text) {
		if i%2 != 0 {
			fmt.Print(string(value))
		}
	}
}
