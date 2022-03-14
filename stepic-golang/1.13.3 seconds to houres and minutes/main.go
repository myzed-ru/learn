package main

import "fmt"

func main() {
	var seconds, hours, minutes int
	fmt.Scan(&seconds)
	hours = seconds / 3600
	minutes = (seconds - hours*3600) / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
