package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0-1.0)

	// string
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	// Boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// largest n digits numbers: 10^n - 1
// 10^1 - 1 = 9
// 10^2 - 1 = 99
// 10^8 - 1 = 99999999
