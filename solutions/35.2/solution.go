package main

import "fmt"

func main() {
	team := map[string]float32{
		"P1": 1.98,
		"P2": 2.05,
		"P3": 1.89,
		"P4": 2.0,
		"P5": 2.11}

	var avg float32 = 0
	for _, height := range team {
		avg += height
	}
	avg /= float32(len(team))
	fmt.Print(avg)

}
