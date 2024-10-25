package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}
func main() {
	radius := 4.0
	area := circleArea(radius)
	fmt.Printf("Circle Area: %f\n", area)
}
