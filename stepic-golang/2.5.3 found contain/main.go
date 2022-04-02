package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	var subtext string
	fmt.Scan(&text, &subtext)
	fmt.Print(strings.Index(text, subtext))
}
