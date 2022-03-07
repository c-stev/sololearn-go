package main

import "fmt"

//create the route() function
func route(cities ...string) {
	for _, city := range cities {
		fmt.Printf("%s->", city)
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	var cities []string
	//take n strings as input and append them to the slice
	var city string
	for i := 0; i < n; i++ {
		fmt.Scan(&city)
		cities = append(cities, city)
	}

	//
	route(cities...)
}
