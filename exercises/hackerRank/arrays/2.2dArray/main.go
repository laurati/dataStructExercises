package main

import (
	"math"
)

// Given a  2D Array, :

// 1 1 1 0 0 0
// 0 1 0 0 0 0
// 1 1 1 0 0 0
// 0 0 0 0 0 0
// 0 0 0 0 0 0
// 0 0 0 0 0 0
// An hourglass in  is a subset of values with indices falling in this pattern in 's graphical representation:

// a b c
//   d
// e f g

// There are  hourglasses in . An hourglass sum is the sum of an hourglass' values.
// Calculate the hourglass sum for every hourglass in , then print the maximum hourglass sum. The array will always be .

// Example

// -9 -9 -9  1 1 1
//  0 -9  0  4 3 2
// -9 -9 -9  1 2 3
//  0  0  8  6 6 0
//  0  0  0 -2 0 0
//  0  0  1  2 4 0
// The  hourglass sums are:

// -63, -34, -9, 12,
// -10,   0, 28, 23,
// -27, -11, -2, 10,
//   9,  17, 25, 18
// The highest hourglass sum is 28

func main() {

}

func hourglassSum(arr [][]int32) int32 {

	var max int32 = math.MinInt32

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			var sum int32 = 0
			sum += arr[row][col]
			sum += arr[row][col+1]
			sum += arr[row][col+2]
			sum += arr[row+1][col+1]
			sum += arr[row+2][col]
			sum += arr[row+2][col+1]
			sum += arr[row+2][col+2]

			if sum > max {
				max = sum
			}
		}
	}
	return max
}
