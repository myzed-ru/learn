package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	allSymbols := []rune(text)
	if unicode.IsUpper(allSymbols[0]) && string(allSymbols[len(allSymbols)-2]) == "." {
		// For stepic: if unicode.IsUpper(allSymbols[0]) && string(allSymbols[len(allSymbols)-1]) == "." {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}

}
