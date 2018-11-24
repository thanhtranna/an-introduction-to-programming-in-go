package main

import (
	"fmt"
	"os"
)

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func testPanic1() {
	panic("PANIC")
	str := recover()
	fmt.Println(str)
}

func testPanic2() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("Panic")
}

func main() {
	defer second()
	second()
	first()

	filename := "test.txt"
	f, _ := os.Open(filename)
	defer f.Close()

	// Will occurs error
	// testPanic1()
	testPanic2()
}
