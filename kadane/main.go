package main

import (
	"fmt"
)

// encontrar a combinação de elementos da lista que resulta na soma máxima

func maxSubArray(nums []int) int {
	maxEndingHere := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	lista := []int{2, 7, 6, 10, -1, -5}
	maxSum := maxSubArray(lista)
	fmt.Println("A soma máxima da sublista é:", maxSum)
}
