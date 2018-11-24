package main

/*
  The Fibonacci sequence is defined as:
  fib(0) = 0,
  fib(1) = 1,
  fib(n) = fib(n-1) + fib(n-2)

  Write a recursive function which can find fib(n).
*/

import "fmt"

func fib(x uint) uint {
	if x <= 1 {
		return x
	}

	return fib(x-1) + fib(x-2)
}

func fibonacci(n uint) uint {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println(fib(10))

	// Other way
	fmt.Println(fibonacci(5))
}
