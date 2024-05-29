package main

import "fmt"

type node struct {
	data int
	next *node
}

// O(1)
func (n *node) insert(x int) *node {
	newNode := &node{
		data: x,
		next: n,
	}
	return newNode
}

// O(1)
func (n *node) delete() {
	n.data = n.next.data
	n.next = n.next.next
}

// O(n)
func (n *node) printList() {
	current := n
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// O(n)
func (n *node) searchValue(x int) bool {
	current := n
	for current != nil {
		if current.data == x {
			return true
		}
		current = current.next
	}
	return false
}

func main() {

	var n *node

	n = n.insert(5)
	n = n.insert(7)
	n = n.insert(9)

	n.printList()

	n.delete()
	fmt.Println("após delete")
	n.printList()

	fmt.Println(n.searchValue(5))
}
