package main

import "fmt"

func main() {

	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	fmt.Println(sockMerchant(9, ar))

}

func sockMerchant(n int32, ar []int32) int32 {

	count := 0
	socks := make(map[int32]int32)

	for _, v := range ar {
		socks[v]++

		if socks[v] > 1 && socks[v]%2 == 0 {
			count++
		}
	}

	return int32(count)

}
