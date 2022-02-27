package main

import "fmt"

func main() {
	var number int
	for fmt.Scan(&number); number <= 100; fmt.Scan(&number) {
		if number >= 10 {
			fmt.Println(number)
		}
	}
}
