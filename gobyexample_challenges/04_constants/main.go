package main

import (
	"fmt"
	"math"
)

func main() {
	// TODO: Declare a constant for Pi and calculate the area of a circle with radius 5
	const Pi = 3.14
	var radius float64 = 5
	area := Pi * math.Pow(radius, 2)
	fmt.Printf("Area of the circle with radius %.2f is %.2f\n", radius, area)
}
