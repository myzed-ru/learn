package main

import (
	"fmt"
)

func main() {
	var number int
	var cows int
	fmt.Scan(&cows)
	number = cows
	if number >= 100 {
		for ; number/100 != 0; number %= 100 {
		}
	}

	switch {
	case number >= 11 && number <= 14:
		fmt.Printf("%d korov", cows)
	case number%10 == 1:
		fmt.Printf("%d korova", cows)
	case number%10 == 2:
		fmt.Printf("%d korovy", cows)
	case number%10 == 3:
		fmt.Printf("%d korovy", cows)
	case number%10 == 4:
		fmt.Printf("%d korovy", cows)
	default:
		fmt.Printf("%d korov", cows)
	}

}
