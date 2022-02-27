package main

import "fmt"

func main() {
	var count, number int
	var summ int = 0
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		fmt.Scan(&number)
		if (number < 100) && (number >= 10) && (number%8 == 0) {
			summ += number
		}
	}
	fmt.Println(summ)
}
