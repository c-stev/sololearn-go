package main

import "fmt"

func main() {
	var f int
	fmt.Scanln(&f)
	//your code goes here
	if f < 0 {
		fmt.Print("Wrong Input")
	} else if f < 20 {
		fmt.Print("Infrasound")
	} else if f <= 20000 {
		fmt.Print("Audible")
	} else {
		fmt.Print("Ultrasound")
	}
}
