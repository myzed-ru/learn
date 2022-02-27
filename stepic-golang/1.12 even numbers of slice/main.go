package main

import "fmt"

func main() {
	var lenArray int
	fmt.Scan(&lenArray)
	myArray := make([]int, lenArray, lenArray)
	for i, _ := range myArray {
		fmt.Scan(&myArray[i])
	}
	for i := 0; i < len(myArray); i += 2 {
		fmt.Printf("%d ", myArray[i])
	}
}
