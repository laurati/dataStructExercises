package main

import "fmt"

func main() {

	a := []int{4, 2, 5, 1, 3}

	fmt.Println(bubbleSort(a))
	// countSwaps(a)

}

func bubbleSort(a []int) []int {

	countSwaps := 0

	for k := 0; k < len(a)-1; k++ {
		swapped := false

		for i := 0; i < len(a)-1-k; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
				countSwaps++
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Println("countSwapps: ", countSwaps)
	return a
}

// func countSwaps(a []int) {
// 	countSwaps := 0

// 	for k := 0; k < len(a)-1; k++ {
// 		swapped := false

// 		for i := 0; i < len(a)-1-k; i++ {
// 			if a[i] > a[i+1] {
// 				a[i], a[i+1] = a[i+1], a[i]
// 				swapped = true
// 				countSwaps++
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	fmt.Printf("Array is sorted in %d swaps.\n", countSwaps)
// 	fmt.Printf("First Element: %d\n", a[0])
// 	fmt.Printf("Last Element: %d\n", a[len(a)-1])

// }
