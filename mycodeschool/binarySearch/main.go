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

	// find count of element in sorted list
	c := []int{1, 1, 3, 3, 5, 5, 5, 5, 5, 9, 9, 11}
	fmt.Println(binarySearchCountElement(c, 2))

	r := []int{11, 12, 15, 18, 2, 5, 6, 8}
	fmt.Println(FindRotationCount(r))

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
		} else if x > a[midIndex] {
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
			result = midIndex
			start = midIndex + 1
		} else if x > a[midIndex] {
			start = midIndex + 1
		} else {
			end = midIndex - 1
		}
	}
	return result
}

func binarySearchCountElement(c []int, x int) int {
	first := binarySearchFirstOccurence(c, x)
	if first == -1 {
		return 0
	}
	last := binarySearchLastOccurence(c, x)

	return (last - first + 1)
}

func FindRotationCount(a []int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		if a[low] <= a[high] {
			return low
		}
		mid := (low + high) / 2
		next := (mid + 1) % (len(a) - 1)
		prev := (mid + len(a) - 1) % (len(a) - 1)

		if a[mid] <= a[next] && a[mid] <= a[prev] {
			return mid
		} else if a[mid] <= a[high] {
			high = mid - 1
		} else if a[mid] >= a[low] {
			low = mid + 1
		}
	}
	return -1
}
