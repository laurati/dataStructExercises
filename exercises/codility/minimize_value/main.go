package main

import (
	"fmt"
	"math"
)

func main() {

	A := []int{3, 1, 2, 4, 3}

	// solução O(n*n)
	// sums1 := 0
	// min := 101000

	// for i := 1; i <= len(A)-1; i++ {
	// 	sums2 := 0
	// 	s1 := A[:i]
	// 	sums1 += s1[i-1]
	// 	s2 := A[i:]
	// 	for _, v := range s2 {
	// 		sums2 += v
	// 	}

	// 	dif := sums1 - sums2
	// 	difAbs := math.Abs(float64(dif))
	// 	difInt := int(difAbs)

	// 	if difInt < min {
	// 		min = difInt
	// 	}
	// }
	// fmt.Println(min)

	// solução O(n)
	totalSum := 0
	for _, num := range A {
		totalSum += num
	}

	leftSum := 0
	minDiff := math.MaxInt64

	for i := 0; i < len(A)-1; i++ {
		leftSum += A[i]
		rightSum := totalSum - leftSum
		diff := int(math.Abs(float64(leftSum - rightSum)))
		if diff < minDiff {
			minDiff = diff
		}
	}
	fmt.Println(minDiff)
}
