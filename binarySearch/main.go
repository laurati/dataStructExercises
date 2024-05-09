package main

import "fmt"

func main() {

	a := []int{2, 4, 6, 8, 10, 11}

	fmt.Println(binarySearch(a, 6))
	fmt.Println(binarySearchRecursive(a, 6, 0, len(a)-1))

	// first and last occurence
	b := []int{2, 4, 10, 10, 10, 18, 20}
	fmt.Println(binarySearchFirstOccurence(b, 10))
	fmt.Println(binarySearchLastOccurence(b, 10))
}

func binarySearch(a []int, x int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		var midIndex int = (start + end) / 2 //outra forma: start + (end - start)
		if a[midIndex] == x {
			return midIndex
		}
		if x > a[midIndex] {
			start = midIndex + 1
		} else {
			end = midIndex - 1
		}
	}
	return -1
}

func binarySearchRecursive(a []int, x, low, high int) int {
	if low > high {
		return -1
	}
	var mid int = (low + high)
	if x == a[mid] {
		return mid
	} else if x < a[mid] {
		return binarySearchRecursive(a, x, low, mid-1)
	} else {
		return binarySearchRecursive(a, x, mid+1, high)
	}
}

func binarySearchFirstOccurence(a []int, x int) int {
	start := 0
	end := len(a) - 1
	result := -1

	for start <= end {
		var midIndex int = (start + end) / 2
		if a[midIndex] == x {
			result, end = midIndex, midIndex-1
		}
		if x > a[midIndex] {
			start = midIndex + 1
		} else {
			end = midIndex - 1
		}
	}
	return result
}

func binarySearchLastOccurence(a []int, x int) int {
	start := 0
	end := len(a) - 1
	result := -1

	for start <= end {
		var midIndex int = (start + end) / 2
		if a[midIndex] == x {
			result, start = midIndex, midIndex+1
		}
		if x > a[midIndex] {
			start = midIndex + 1
		} else {
			end = midIndex - 1
		}
	}
	return result
}
