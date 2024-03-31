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

func (l *linkedList) prepend(n *node) {
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

func (l linkedList) printListData() {
	toPrint := l.head
	for toPrint != nil {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Printf("\n")
}

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

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.printListData()

	myList.deleteWithValue(100)
	myList.deleteWithValue(2)
	myList.printListData()

	emptyList := linkedList{}
	emptyList.deleteWithValue(10)
}
