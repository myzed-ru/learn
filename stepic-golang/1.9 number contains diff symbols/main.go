package main

import "fmt"

func main() {
	var number, first, second, third int

	fmt.Scan(&number)
	first = number / 100
	second = number % 100 / 10
	third = number % 100 % 10
	if first != second && first != third && second != third {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
