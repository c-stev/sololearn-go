package main

import "fmt"

func main() {
	//your code goes here
	var temp float64
	fmt.Scan(&temp)
	if temp < 99.5 {
		fmt.Print("Allowed")
	} else {
		fmt.Print("Fever")
	}
}
