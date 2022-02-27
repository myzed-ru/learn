package main

import "fmt"

func main() {
	var ticket string
	fmt.Scan(&ticket)
	if (int(ticket[0]) + int(ticket[1]) + int(ticket[2])) == (int(ticket[3]) + int(ticket[4]) + int(ticket[5])) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
