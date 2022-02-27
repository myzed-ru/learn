package main

import "fmt"

func main() {
	var number float64
	fmt.Scan(&number)
	switch {
	case number <= 0:
		fmt.Printf("число %2.2f не подходит", number)
	case number > 10000:
		fmt.Printf("%e", number)
	default:
		fmt.Printf("%.4f", number*number)
	}
}
