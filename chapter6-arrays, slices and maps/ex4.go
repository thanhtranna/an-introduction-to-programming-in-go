package main

import (
	"fmt"
)

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var smallest int
	smallest = x[0]
	length := len(x)
	fmt.Println(length)
	for i := 1; i < length; i++ {
		fmt.Println(x[i])
		if x[i] < smallest {
			smallest = x[i]
		}
	}
	fmt.Println("Smallest:", smallest)

	// Other way
	val := x[0]
	for _, value := range x {
		if value < val {
			val = value
		}
	}
	fmt.Println("Smallest value:", val)
}
