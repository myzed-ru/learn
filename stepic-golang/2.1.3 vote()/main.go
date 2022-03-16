package main

import "fmt"

func main() {
	fmt.Print(vote(1, 0, 0))
}

func vote(x int, y int, z int) int {
	switch {
	case x == y:
		return x
	case y == z:
		return y
	}
	return z
}
