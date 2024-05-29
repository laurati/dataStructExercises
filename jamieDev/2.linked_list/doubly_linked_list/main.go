package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

// Insert at the beginning
func (l *linkedList) insert(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		second := l.head
		l.head = n
		n.next = second
		second.prev = n
	}
	l.length++
}

// Insert at the end
func (l *linkedList) insertAtEnd(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for toPrint != nil {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Printf("\n")
}

// Print list from tail to head
func (l linkedList) PrintBackward() {
	current := l.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Printf("\n")
}

// como tem apenas o valor, a complexidade fica O(n)
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		l.length--
		return
	}
	current := l.head
	for current != nil && current.data != value {
		current = current.next
	}
	if current == nil {
		return
	}
	if current == l.tail {
		l.tail = current.prev
	}
	if current.prev != nil {
		current.prev.next = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	}
	l.length--
}

// Remove um nó específico da lista
// se tiver o node,a complexidade fica O(1)
func (l *linkedList) Remove(node *node) {
	if node == nil {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}

	myList.insertAtEnd(node1)
	myList.insertAtEnd(node2)
	myList.insertAtEnd(node3)
	myList.insertAtEnd(node4)
	myList.insertAtEnd(node5)
	myList.insertAtEnd(node6)

	myList.printListData()
	fmt.Println("=========")
	myList.PrintBackward()

	myList.Remove(node4)
	myList.printListData()

	// myList.deleteWithValue(100)
	// myList.deleteWithValue(2)
	// myList.printListData()

	// emptyList := linkedList{}
	// emptyList.deleteWithValue(10)
}
