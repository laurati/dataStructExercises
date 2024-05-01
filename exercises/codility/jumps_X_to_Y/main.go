package main

import (
	"fmt"
	"math"
)

func main() {

	X := 10
	Y := 85
	D := 30

	a := 0
	count := 0

	for {
		X += D

		a = X
		count++
		if a >= Y {
			break
		}

	}
	fmt.Println(a, count)

	// forma mais eficiente
	distance := Y - X
	jumps := math.Ceil(float64(distance) / float64(D))
	fmt.Println(jumps)

}
