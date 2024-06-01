package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Função para trocar dois elementos no slice
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Função de partição normal
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

// Função de partição randomizada
func randomizedPartition(arr []int, low, high int) int {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(high-low+1) + low
	swap(arr, randomIndex, high)
	return partition(arr, low, high)
}

// Função principal do Quick Sort
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := randomizedPartition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)
	quickSort(arr, 0, n-1)
	fmt.Println("Sorted array:", arr)
}
