package main

import "fmt"

func main() {
	var number int
	var root int = 0
	fmt.Scan(&number)
	for {
		for number != 0 {
			root += number % 10
			number = number / 10
		}
		if root/10 == 0 {
			break
		}
		number = root
		root = 0
	}
	fmt.Print(root)

}
