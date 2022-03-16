package main

import (
	"fmt"
)

func main() {
	var digits, digit string
	fmt.Scan(&digits, &digit)
	for i := 0; i < len(digits); i++ {
		if digit != string(digits[i]) {
			fmt.Print(string(digits[i]))
		}
	}
}
