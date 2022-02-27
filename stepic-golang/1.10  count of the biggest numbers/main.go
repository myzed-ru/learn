package main

import "fmt"

func main() {
	var max, number, count int
	max = number
	number = 1
	count = 0
	for number != 0 {
		fmt.Scan(&number)
		switch {
		case number > max:
			{
				max = number
				count = 1
			}
		case number == max:
			{
				count++
			}
		}
	}
	fmt.Println(count)

}
