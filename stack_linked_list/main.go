package main

import "fmt"

type node struct {
	v    int
	next *node
}

type Stack struct {
	head *node
}

func (s *Stack) push(n *node) {
	if s.head == nil {
		s.head = n
		return
	}
	second := s.head
	s.head = n
	s.head.next = second
}

func (s *Stack) pop() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
}

func (s Stack) print() {
	for s.head != nil {
		fmt.Println(s.head.v)
		s.head = s.head.next
	}

}

func main() {

	list := Stack{}

	n1 := node{v: 3}
	n2 := node{v: 2}
	n3 := node{v: 1}

	list.push(&n1)
	list.push(&n2)
	list.push(&n3)

	list.print()

	list.pop()

	fmt.Println("after pop")
	list.print()

}
