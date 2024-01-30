package main

import "fmt"

func main() {
	// ex 1 - reverse
	// a := []int32{4, 3, 2, 1}

	// var b []int32

	// for i := len(a) - 1; i >= 0; i-- {
	// 	b = append(b, a[i])
	// }
	// fmt.Println(b)

	// ex 2 - hourglass
	// arr := [][]int32{

	// 	// 19
	// 	// {1, 1, 1, 0, 0, 0},
	// 	// {0, 1, 0, 0, 0, 0},
	// 	// {1, 1, 1, 0, 0, 0},
	// 	// {0, 0, 2, 4, 4, 0},
	// 	// {0, 0, 0, 2, 0, 0},
	// 	// {0, 0, 1, 2, 4, 0},

	// 	// -6
	// 	// {-1, -1, 0, -9, -2, -2},
	// 	// {-2, -1, -6, -8, -2, -5},
	// 	// {-1, -1, -1, -2, -3, -4},
	// 	// {-1, -9, -2, -4, -4, -5},
	// 	// {-7, -3, -3, -2, -9, -9},
	// 	// {-1, -3, -1, -2, -4, -5},

	// 	// -19
	// 	{0, -4, -6, 0, -7, -6},
	// 	{-1, -2, -6, -8, -3, -1},
	// 	{-8, -4, -2, -8, -8, -6},
	// 	{-3, -1, -2, -5, -7, -4},
	// 	{-3, -5, -3, -6, -6, -6},
	// 	{-3, -6, 0, -8, -6, -7},
	// }

	// hourglasses := make([]int32, 0)

	// maxHourglass := math.MinInt32
	// maxHourglass32 := int32(maxHourglass)

	// for row := 0; row <= 3; row++ {
	// 	for col := 0; col <= 3; col++ {
	// 		var sum int32
	// 		sum = 0

	// 		sum += arr[row][col]
	// 		sum += arr[row][col+1]
	// 		sum += arr[row][col+2]
	// 		sum += arr[row+1][col+1]
	// 		sum += arr[row+2][col]
	// 		sum += arr[row+2][col+1]
	// 		sum += arr[row+2][col+2]

	// 		hourglasses = append(hourglasses, sum)

	// 		if sum > maxHourglass32 {
	// 			maxHourglass32 = sum
	// 		}
	// 	}
	// }

	// // Imprimir os valores calculados
	// fmt.Println("Hourglasses:", hourglasses)
	// fmt.Println("Valor máximo de Hourglasses:", maxHourglass32)

	// ex 3 - dynamicArray
	queries := [][]int32{
		{2, 5},
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

	n := queries[0][0]
	Q := queries[0][1]

	var answerArr []int32

	arr := make([][]int32, n)
	var lastAnswer int32
	lastAnswer = 0

	for i := 0; i < int(Q); i++ {

		queryType := queries[i+1][0]
		x := queries[i+1][1]
		y := queries[i+1][2]

		idx := (x ^ lastAnswer) % n
		seq := arr[idx]

		if queryType == 1 {
			arr[idx] = append(seq, int32(y))
		} else {
			colIndex := y % int32(len(seq))
			lastAnswer = int32(seq[colIndex])
			answerArr = append(answerArr, lastAnswer)
		}
	}
	fmt.Println(answerArr)

}
