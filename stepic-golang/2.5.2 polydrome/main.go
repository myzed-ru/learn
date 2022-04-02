package main

import (
	"fmt"
)

func isPolydrome(s string) bool {
	allSymbols := []rune(s)
	for i := 0; i < len(allSymbols)/2; i++ {
		if allSymbols[i] != allSymbols[len(allSymbols)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var text string
	fmt.Scan(&text)
	if isPolydrome(text) {
		fmt.Print("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
