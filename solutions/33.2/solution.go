package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	//your code goes here
	arr := make([]int, n)
	var in int
	for i := 0; i < n; i++ {
		fmt.Scanln(&in)
		arr[i] = in
	}

	fmt.Print(arr)

}
