package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {

	binaryStr := strconv.FormatInt(int64(N), 2)

	maxGap := 0
	currentGap := 0
	startCounting := false

	for _, bit := range binaryStr {
		if bit == '1' {

			if !startCounting {
				startCounting = true
			} else {

				if currentGap > maxGap {
					maxGap = currentGap
				}
				currentGap = 0
			}
		} else if startCounting {

			currentGap++
		}
	}

	return maxGap
}

func main() {
	fmt.Println(Solution(1041)) // Output: 5
	fmt.Println(Solution(32))   // Output: 0
}
