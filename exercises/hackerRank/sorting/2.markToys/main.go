package main

import (
	"fmt"
	"sort"
)

func main() {

	// prices := []int32{1, 2, 3, 4}

	prices := []int32{1, 12, 5, 111, 200, 1000, 10}

	fmt.Println(maximumToys(prices, 50))

}

func maximumToys(prices []int32, k int32) int32 {

	sort.Slice(prices, func(a, b int) bool {
		return prices[a] < prices[b]
	})

	var sum int32 = 0
	var itens int32 = 0
	for _, v := range prices {
		sum += v
		if sum >= k {
			break
		}
		itens++
	}
	fmt.Println(prices)
	return itens

}

// Time limit exceeded - bubble sort: O(n^2)
// func maximumToys(prices []int32, k int32) int32 {

// 	for k := 0; k < len(prices)-1; k++ {
// 		swapped := false
// 		for i := 0; i < len(prices)-1-k; i++ {
// 			if prices[i] > prices[i+1] {
// 				prices[i], prices[i+1] = prices[i+1], prices[i]
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}

// 	var sum int32 = 0
// 	var itens int32 = 0
// 	for _, v := range prices {
// 		sum += v
// 		if sum >= k {
// 			break
// 		}
// 		itens++
// 	}
// 	return itens

// }
