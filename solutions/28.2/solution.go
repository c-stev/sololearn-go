package main

import "fmt"

type Timer struct {
	id      string
	seconds int
}

func reset(t *Timer) {
	t.seconds = 0
}

func main() {
	var x int
	fmt.Scanln(&x)
	t := Timer{"ABC", x}

	reset(&t)
	fmt.Println(t)
}
