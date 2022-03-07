package main

import "fmt"

type Cart struct {
	prices []float32
}

func (x Cart) show() {
	var sum float32 = 0.0
	//calculate the sum of all prices in the Cart
	for _, value := range x.prices {
		sum += value
	}
	fmt.Println(sum)
}

func main() {
	c := Cart{}
	var n int
	fmt.Scanln(&n)

	// take n inputs and add them to the Cart
	var price float32
	arr := make([]float32, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&price)
		arr[i] = price
	}

	c.prices = arr

	//call the show() method of the Cart
	c.show()

}
