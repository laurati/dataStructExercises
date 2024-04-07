package main

import (
	"fmt"
)

func removeDuplicatesL(nums []int) int {
	m := map[int]struct{}{}
	i := 0

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			i++
		}
	}
	fmt.Println(nums)
	return i
}

func removeDuplicates(nums []int) int {
	prev := nums[0]
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[j] = nums[i]
			j++
		}
		prev = nums[i]
	}
	return j
}

func main() {

	nums := []int{1, 1, 2}

	unique := removeDuplicatesL(nums)

	fmt.Println(unique)
}
