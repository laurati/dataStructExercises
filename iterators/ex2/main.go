package main

import (
	"fmt"
	"iter"
)

// Define o tipo Set com a restrição `comparable`
type Set[V comparable] map[V]struct{}

// Implementa o método All para o Set
func (s Set[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func (s Set[V]) Filtered(predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if predicate(v) && !yield(v) {
				return
			}
		}
	}
}

func Map[V any, R any](seq iter.Seq[V], mapper func(V) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

func main() {
	// Cria um conjunto e itera sobre os elementos
	s := Set[int]{1: {}, 2: {}, 3: {}, 4: {}}
	for v := range s.All() {
		fmt.Println(v)
	}

	fmt.Println("»»»»»»»»»")

	for v := range s.Filtered(func(v int) bool { return v%2 == 0 }) {
		fmt.Println(v)
	}

	fmt.Println("»»»»»»»»»")

	// Mapeia os valores para suas versões duplicadas
	mapped := Map(s.All(), func(v int) int {
		return v * 2
	})

	for v := range mapped {
		fmt.Println(v)
	}
}
