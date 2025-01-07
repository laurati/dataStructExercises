package main

import (
	"fmt"
	"iter"
)

// PrintAll consome o iterador e imprime os elementos
func PrintAll[V any](seq iter.Seq[V]) {
	for v := range seq {
		fmt.Println(v)
	}
}

// Gerador de números como um iterador
func Numbers(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	// Imprime os números de 1 a 5
	PrintAll(Numbers(10))
}
