package main

import "fmt"

func minimumSwaps(arr []int) int {
	n := len(arr)
	visited := make([]bool, n)
	swaps := 0

	for i := 0; i < n; i++ {
		if !visited[i] && arr[i] != i+1 {
			cycleLength := 0
			x := i

			for !visited[x] {
				visited[x] = true
				x = arr[x] - 1
				cycleLength++
			}

			if cycleLength > 0 {
				swaps += cycleLength - 1
			}
		}
	}

	return swaps
}

func main() {
	arr1 := []int{7, 1, 3, 2, 4, 5, 6}
	fmt.Println(minimumSwaps(arr1)) // Output: 5

	// arr2 := []int{4, 3, 1, 2}
	// fmt.Println(minimumSwaps(arr2)) // Output: 3

	// arr3 := []int{2, 3, 4, 1, 5}
	// fmt.Println(minimumSwaps(arr3)) // Output: 3

	// arr4 := []int{1, 3, 5, 2, 4, 6, 7}
	// fmt.Println(minimumSwaps(arr4)) // Output: 3
}
