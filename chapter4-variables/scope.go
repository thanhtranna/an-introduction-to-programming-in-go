package main

import "fmt"

var x string = "Hello World2"

func main() {
	var y string = "Hello World1"
	fmt.Println(x)

}

// undefined: y error
func f() {
	fmt.Println(x)
	fmt.Println(y)
}
