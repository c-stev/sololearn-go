package main

import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	var last string

	fmt.Scanln(&last, "last result")

	total := 0

	for _, i := range results {
		switch i {
		case "w":
			total += 3
		case "d":
			total += 1
		case "l":
			total += 0
		}
	}

	switch last {
	case "w":
		total += 3
	case "d":
		total += 1
	case "l":
		total += 0
	}

	fmt.Println(total)

}
