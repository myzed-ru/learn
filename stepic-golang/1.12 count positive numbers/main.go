package main

import "fmt"

func main() {
	var lenArray int
	var count int = 0
	fmt.Scan(&lenArray)
	myArray := make([]int, lenArray, lenArray)
	for i, _ := range myArray {
		fmt.Scan(&myArray[i])
	}
	for _, element := range myArray {
		if element > 0 {
			count++
		}
	}
	fmt.Print(count)
}
