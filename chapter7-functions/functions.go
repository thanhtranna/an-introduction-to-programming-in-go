package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	someOtherName := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(someOtherName))

	fmt.Println(add(1, 2, 3))

	// closure
	addc := func(x, y int) int {
		return x + y
	}

	fmt.Println(addc(1, 1))

	u := 0
	increment := func() int {
		u++
		return u
	}
	fmt.Println(increment())
	fmt.Println(increment(), u)

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(factorial(4))
}
