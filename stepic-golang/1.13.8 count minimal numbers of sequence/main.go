package main

import "fmt"

func main() {
	var lenArray int
	var min_count int = 0
	var minimanl int = 2147483647
	fmt.Scan(&lenArray)
	myArray := make([]int, lenArray, lenArray)
	for i, _ := range myArray {
		fmt.Scan(&myArray[i])
		if myArray[i] < minimanl {
			minimanl = myArray[i]
		}
	}
	for _, value := range myArray {
		if value == minimanl {
			min_count++
		}
	}
	fmt.Print(min_count)
}
