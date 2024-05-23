package main

import "fmt"

type Stack struct {
	array []int
}

func (s *Stack) push(v ...int) {
	s.array = append(s.array, v...)
}

func (s *Stack) pop() {
	// [a:b] - a é incluso e b é excluso
	s.array = s.array[:len(s.array)-1]
}

func (s *Stack) stackIsNull() bool {
	return len(s.array) == 0
}

func (s *Stack) destroyStack() {
	s.array = s.array[len(s.array)-1 : len(s.array)-1]
}

func main() {

	stack := Stack{}

	stack.push(3, 4, 9, 2)
	fmt.Println(&stack)

	stack.pop()
	fmt.Println(&stack)

	fmt.Println(stack.stackIsNull())

	stack.destroyStack()
	fmt.Println(&stack)

	fmt.Println(stack.stackIsNull())

}
