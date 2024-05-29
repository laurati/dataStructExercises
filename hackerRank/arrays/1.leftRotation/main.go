package main

import "fmt"

func main() {

	a := []int32{1, 2, 3, 4, 5}

	var d int32 = 4

	for d != 0 {

		primeiro := a[0:1]
		a = a[1:]
		a = append(a, primeiro...)
		d--

	}

	fmt.Println(a)

}
