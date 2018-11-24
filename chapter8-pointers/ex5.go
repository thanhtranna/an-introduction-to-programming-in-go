package main

import (
	"fmt"
)

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x, y := 1, 2
	fmt.Printf("Before swap: x=%d ; y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x=%d ; y=%d\n", x, y)
}
