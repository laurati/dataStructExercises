package main

import "fmt"

func main() {

	A := []int{2, 3, 1, 5}

	n := len(A)
	expectedSum := (n + 1) * (n + 2) / 2
	actualSum := 0

	for _, num := range A {
		actualSum += num
	}

	missingElement := expectedSum - actualSum

	fmt.Println(expectedSum, actualSum, missingElement)
}
