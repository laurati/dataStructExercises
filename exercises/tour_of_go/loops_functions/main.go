package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	previousZ := z
	threshold := 0.000001

	for {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("z = %v\n", z)
		if math.Abs(z-previousZ) < threshold {
			break
		}
		previousZ = z
	}
	return z
}

func main() {
	x := 4.0
	fmt.Println(Sqrt(x))
}
