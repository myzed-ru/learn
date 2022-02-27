package main

import "fmt"

func main() {
	var x, p, y, result float32
	var year int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	year = 0
	result = x
	for {
		result += result / 100 * p
		year++
		if result >= y {
			break
		}
	}
	fmt.Println(year)
}
