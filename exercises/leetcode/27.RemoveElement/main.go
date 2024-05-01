package main

import "fmt"

func countElementDifferentThanVal(nums []int, val int) int {
	i := 0

	for _, v := range nums {
		if v != val {
			// nums[i] = v
			i += 1
		}
	}
	return i

}

func removeElement(nums []int, val int) {

	c := []int{}

	for _, v := range nums {
		if v != val {
			c = append(c, v)
		}
	}
	fmt.Println(c)
}

func main() {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(countElementDifferentThanVal(nums, 2))

	removeElement(nums, 2)
}
