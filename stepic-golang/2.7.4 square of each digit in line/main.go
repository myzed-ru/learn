package main

import (
	"fmt"
	"strconv"
)

func main() {

	var myString string

	_, err := fmt.Scan(&myString)
	if err != nil {
		panic("var myString has invalid value")
	}

	for _, value := range myString {
		intVar, _ := strconv.Atoi(string(value))
		fmt.Print(intVar * intVar)
	}

}
