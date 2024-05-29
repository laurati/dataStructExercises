package main

import "fmt"

func main() {

	count := 0

	// arr := []int{0, 1, 0, 0, 0, 1, 0} //3
	// arr := []int{0, 0, 1, 0, 0, 1, 0} //4
	// arr := []int{0, 0, 0, 0, 1, 0} //3
	arr := []int{0, 0, 0, 1, 0, 0} //3
	// arr := []int{0, 0, 1, 0, 0, 1, 0} //4

	i := 0

	for i < len(arr)-1 {

		if i+2 <= len(arr)-1 && arr[i+2] == 0 {
			count++
			i = i + 2
		} else {
			count++
			i++
		}

	}

	fmt.Println(count)

}
