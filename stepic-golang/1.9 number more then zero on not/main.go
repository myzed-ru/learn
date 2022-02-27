package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	switch {
	case 0 < number:
		fmt.Print("Число положительное")
	case 0 > number:
		fmt.Print("Число отрицательное")
	default:
		fmt.Print("Ноль")
	}
}
