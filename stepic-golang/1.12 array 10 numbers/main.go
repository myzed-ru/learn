package main

import (
	"fmt"
)

func main() {
	var workArray [10]uint8
	var change1, change2 uint8
	for i, _ := range workArray {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&change1, &change2)
		workArray[change1], workArray[change2] = workArray[change2], workArray[change1]
	}
	output := fmt.Sprint(workArray)
	fmt.Print(output[1:len(output)-1] + " ")

}
