package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Yaya", "bb"},     // города с населением 10-99 тыс. человек
		100:  []string{"aa", "cc", "dd"}, // города с населением 100-999 тыс. человек
		1000: []string{"bb", "nn"},       // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Nsk":  10000,
		"Bsk":  100,
		"cc":   150,
		"bb":   1500,
		"Yaya": 50,
	}

	//scan:=0
	cityFound := true
	for name, population := range cityPopulation {
		if population < 100 || population >= 1000 {
			delete(cityPopulation, name)
		} else {
			cityFound = false
			for _, cityName := range groupCity[100] {
				if cityName == name {
					cityFound = true
				}
			}
			if !cityFound {
				delete(cityPopulation, name)
			}
		}
	}
	fmt.Print(cityPopulation)

}
