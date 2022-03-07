package main

import "fmt"

func main() {
	menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

	//your code goes here
	var input int
	fmt.Scan(&input)

	if len(menu) > input {
		fmt.Print(menu[input])
	} else {
		fmt.Print("Invalid choice")
	}

}
