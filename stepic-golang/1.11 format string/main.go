package main

import "fmt"

func main() {
	var my_string string
	my_string = "This string is on\n" +
		"multiple lines\n" +
		"within three single\n" +
		"quotes on either side."
	fmt.Println(my_string)
	fmt.Println(`Sammy says,\"The balloon\'s color is red.\"`)
}
