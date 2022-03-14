package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Print(number/100 + number/10%10 + number%100%10)
}
