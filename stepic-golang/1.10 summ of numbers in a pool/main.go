package main

import "fmt"

func main() {
	var i, start, end, summ int
	fmt.Scan(&start)
	fmt.Scan(&end)
	summ = 0
	for i = start; i <= end; i++ {
		summ += i
	}
	fmt.Println(summ)
}
