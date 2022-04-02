package main

import "fmt"

func main() {
	fSemyon := make([]string, 3)
	fSemyon = append(fSemyon, "Valera", "Tanya", "Dima")
	fmt.Print(fSemyon[3])
	friends := map[string][]string{
		"Kostya": nil,
	}
	v := friends["Kostya"]
	fmt.Print(value)
}
