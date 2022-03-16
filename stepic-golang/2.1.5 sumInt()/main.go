package main

import "fmt"

func main() {
	fmt.Print(sumInt(1, 0))

}
func sumInt(args ...int) (int, int) {
	summa := 0
	for _, arg := range args {
		summa += arg
	}
	return len(args), summa
}
