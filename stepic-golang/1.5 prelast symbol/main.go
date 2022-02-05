package main

import "fmt"

func main() {
	//var natural string
	//fmt.Scan(&natural)
	//fmt.Print(natural[len(natural)-2:len(natural)-1])
	var number int
	fmt.Scan(&number)
	fmt.Print(number % 100 / 10)
}
