package main

/*
 Write another program that converts from feet into meters. (1 ft = 0.3048 m)
*/

import "fmt"

func feetToMeter(feet float64) float64 {
	return feet * 0.3048
}

func meterToFeet(meter float64) float64 {
	return 1.0 / 0.3048 * meter
}

func main() {
	c := 42.22
	f := 12.868656
	fmt.Printf("%f meters equals %f feets\n", c, feetToMeter(c))
	fmt.Printf("%f feets equals %f meters\n", f, meterToFeet(f))
}
