package main

import "fmt"

func main() {

	A := []int{4, 1, 3, 2}

	fmt.Println(Solution(A))

}
func Solution(A []int) int {
	m := make(map[int]bool)

	if len(A) == 0 {
		return 0
	}
	for _, v := range A {
		if v <= 0 || v > len(A) || m[v] {
			return 0
		}
		m[v] = true
	}
	for i := 1; i <= len(A); i++ {
		if !m[i] {
			return 0
		}
	}
	return 1
}
