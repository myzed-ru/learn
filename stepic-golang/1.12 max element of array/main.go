package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	for _, element := range array {
		if element > a {
			a = element
		}
	}
	fmt.Print(a)
}
