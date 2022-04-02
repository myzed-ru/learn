package main

import "fmt"

func main() {
	var inputVar int
	var results = map[int]int{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&inputVar)
		if value, ok := results[inputVar]; ok {
			fmt.Printf("%d ", value)
		} else {
			results[inputVar] = work(inputVar)
			fmt.Printf("%d ", results[inputVar])
		}
	}

}

func work(x int) int {
	//fmt.Print(x)
	return x * 2
}
