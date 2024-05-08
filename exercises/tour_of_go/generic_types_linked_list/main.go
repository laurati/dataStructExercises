package main

import "fmt"

// In addition to generic functions, Go also supports generic types.
// A type can be parameterized with a type parameter,
// which could be useful for implementing generic data structures.
// This example demonstrates a simple type declaration for a singly-linked list holding any type of value.

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Append adds a new element to the end of the list.
func (l *List[T]) Append(val T) {
	if l.next == nil {
		l.next = &List[T]{val: val}
	} else {
		l.next.Append(val)
	}
}

// Print prints the contents of the list.
func (l *List[T]) Print() {
	fmt.Printf("%v ", l.val)
	if l.next != nil {
		l.next.Print()
	}
}

func main() {
	list := &List[int]{val: 1}
	list.Append(2)
	list.Append(3)

	list.Print() // Output: 1 2 3
}
