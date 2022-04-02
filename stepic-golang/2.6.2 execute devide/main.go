package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

func main() {

	var a, b, result int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	result, err = divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	fmt.Println(string(result))
}
