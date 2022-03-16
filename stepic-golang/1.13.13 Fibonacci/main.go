package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	var previous int = 1
	var last int = 1
	var fibonnaci int = 1
	var i int = 1
	for ; fibonnaci <= number; i++ {
		fibonnaci += previous
		previous = last
		last = fibonnaci
	}
	if previous == number {
		fmt.Print(i)
	} else {
		fmt.Print("-1")
	}

}
