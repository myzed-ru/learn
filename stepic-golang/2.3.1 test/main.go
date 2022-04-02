package main

import "fmt"

var a, b int

func main() {
	a, b = 2, 3
	test(&a, &b)
}
func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Printf("%d %d", *x1, *x2)
}
