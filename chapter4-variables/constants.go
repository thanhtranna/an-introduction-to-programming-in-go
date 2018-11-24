package main

import "fmt"

func main() {
	const x string = "Hello World!"

	//cannot assign to x
	x = "This is a test content"
	fmt.Println(x)
}
