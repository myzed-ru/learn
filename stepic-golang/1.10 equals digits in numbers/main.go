package main

import "fmt"

func main() {
	var first, second, result string
	fmt.Scan(&first, &second)
	result = ""
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				result += string(first[i]) + " "
			}
		}
	}
	fmt.Print(result[0 : len(result)-1])
}
