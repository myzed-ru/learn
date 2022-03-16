package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Scan(&number)
	for degree := 1; degree <= number; degree *= 2 {
		fmt.Print(degree, " ")
	}

}
