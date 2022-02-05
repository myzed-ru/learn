package main

import "fmt"

func main() {
	var natural string
	fmt.Scan(&natural)
	fmt.Print(natural[len(natural)-1:])
}
