package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) append(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList) appendInOrder(n *node) {

	current := l.head

	if l.head == nil || n.data <= l.head.data {
		l.head = n
		l.head.next = current
		l.length++
		return
	}
	for current.next != nil && current.next.data < n.data {
		current = current.next
	}
	current.next = n

	l.length++
}

func (l LinkedList) print() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func (l LinkedList) search(v int) bool {
	toPrint := l.head
	for l.length != 0 {
		if toPrint.data == v {
			return true
		}
		toPrint = toPrint.next
		l.length--
	}
	return false
}

func (l *LinkedList) delete(v int) {

	if l.length == 0 {
		fmt.Println("empty list")
		return
	}

	if l.head.data == v {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != v {
		if previousToDelete.next.next == nil {
			fmt.Println("value to delete was not found")
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l *LinkedList) destroyList() {
	l.head = nil
	l.length = 0
}

func compareLists(l1, l2 *LinkedList) bool {
	current1 := l1.head
	current2 := l2.head

	for current1 != nil && current2 != nil {
		if current1.data != current2.data {
			return false
		}
		current1 = current1.next
		current2 = current2.next
	}

	if current1 != nil || current2 != nil {
		return false
	}
	return true
}

func main() {
	node1 := node{data: 2}
	node2 := node{data: 3}
	node3 := node{data: 1}
	node4 := node{data: 4}

	myList := LinkedList{}

	myList.appendInOrder(&node1)
	myList.appendInOrder(&node2)
	myList.appendInOrder(&node3)
	myList.appendInOrder(&node4)

	node5 := node{data: 3}
	node6 := node{data: 2}
	node7 := node{data: 1}

	secondList := LinkedList{}

	secondList.append(&node5)
	secondList.append(&node6)
	secondList.append(&node7)

	myList.print()
	fmt.Println("=============================")
	secondList.print()

	fmt.Println("compare: ", compareLists(&myList, &secondList))

	fmt.Println(secondList.search(2))

	secondList.delete(2)
	secondList.print()

	fmt.Println(secondList.search(2))

	secondList.destroyList()
	secondList.print()
}
