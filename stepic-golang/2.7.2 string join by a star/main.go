package main

import (
	"fmt"
	"strings"
)

func main() {

	var myString string

	_, err := fmt.Scan(&myString)
	if err != nil {
		panic("var myString has invalid value")
	}

	fmt.Print(strings.Join(strings.Split(myString, ""), "*"))

}
