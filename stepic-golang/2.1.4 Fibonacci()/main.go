package main

import "fmt"

func main() {
	fmt.Print(fibonacci(4))

}
func fibonacci(n int) int {

	var previous int = 1
	var last int = 1
	var row int = 1
	var i int = 1
	for ; i <= n-2; i++ {
		row += previous
		previous = last
		last = row
	}
	return row
}
