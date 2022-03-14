package main

import "fmt"

func main() {
	var n, number int
	var count int = 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&number)
		if number == 0 {
			count++
		}
	}
	fmt.Print(count)
}
