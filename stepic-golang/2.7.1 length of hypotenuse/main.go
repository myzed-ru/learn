package main

import (
	"fmt"
	"math"
)

func hypotenuse(a int, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {

	var a, b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		panic("vars a or b has invalid value")
	}

	fmt.Println(hypotenuse(a, b))
}
