package main

import "fmt"

func main() {
	var degree, hour float32
	var minute int
	fmt.Scan(&degree)
	hour = 12.0 / 360.0 * degree
	minute = int(degree) % 30 * 2
	fmt.Println("It is", int(hour), "hours", int(minute), "minutes.")
}
