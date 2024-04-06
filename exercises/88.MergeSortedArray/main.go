package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) > m {
		nums1 = nums1[:m]
	}
	if len(nums2) > n {
		nums2 = nums2[:n]
	}

	nums1 = append(nums1, nums2...)

	// sort.Ints(nums1)

	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {

			if nums1[j] < nums1[i] {
				menor := nums1[j]
				nums1[j] = nums1[i]
				nums1[i] = menor

			}
		}
	}

	fmt.Println(nums1)
}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)

}
