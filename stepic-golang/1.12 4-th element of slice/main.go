package main

import "fmt"

func main() {
	var lenArray int
	fmt.Scan(&lenArray)
	myArray := make([]int, lenArray, lenArray)
	for i, _ := range myArray {
		fmt.Scan(&myArray[i])
	}
	fmt.Print(myArray[3])
}
