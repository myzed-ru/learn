package main

import (
	"fmt"
)

func main() {

	var myString string

	_, err := fmt.Scan(&myString)
	if err != nil {
		panic("var myString has invalid value")
	}

	allDigitals := []rune(myString)
	max := allDigitals[0]
	for _, value := range allDigitals {
		if value > max {
			max = value
		}
	}
	fmt.Println(string(max))

}
