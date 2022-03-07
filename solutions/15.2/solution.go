package main

import (
	"fmt"
)

func main() {
	var years int
	fmt.Scanln(&years)

	//your code goes here
	output := 7
	for i := 0; i < years; i++ {
		output *= 2
	}
	fmt.Print(output)

}
