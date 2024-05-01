package main

import "fmt"

func main() {

	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(nums))

}

func majorityElement(nums []int) int {

	m := make(map[int]int)
	maxFrequency := 0
	numMaisRepete := 0

	for _, v := range nums {
		m[v]++
		if m[v] > maxFrequency {
			maxFrequency = m[v]
			numMaisRepete = v
		}
	}
	return numMaisRepete
}
