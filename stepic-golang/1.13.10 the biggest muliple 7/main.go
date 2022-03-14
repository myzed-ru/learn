package main

import (
	"fmt"
)

func main() {
	var min, max, number int

	fmt.Scan(&min, &max)
	if min == 0 || max == 0 {
		fmt.Print(0)
	} else {
		for ; ; max-- {
			if max%7 == 0 {
				number = max / 7
				fmt.Print(max)
				break
			}
			if max <= min {
				break
			}
		}

		if number == 0 {
			fmt.Print("NO")
		}
	}

}
