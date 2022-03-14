package main

import "fmt"

func main() {
	var number string
	fmt.Scan(&number)
	fmt.Print(string(number[2]), string(number[1]), string(number[0]))
}
