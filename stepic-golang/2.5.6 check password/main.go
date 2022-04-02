package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)
	allSymbols := []rune(password)
	if len(allSymbols) < 5 {
		fmt.Print("Wrong password")
		return
	}
	for _, symbol := range allSymbols {
		if !(unicode.Is(unicode.Latin, symbol) || unicode.IsDigit(symbol)) {
			fmt.Print("Wrong password")
			return
		}
	}
	fmt.Print("Ok")
}
