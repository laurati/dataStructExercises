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

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
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

func main() {
	node1 := node{data: 2}
	node2 := node{data: 1}
	node3 := node{data: 3}

	myList := LinkedList{}

	myList.prepend(&node1)
	myList.prepend(&node2)
	myList.prepend(&node3)

	myList.print()
	fmt.Println(myList)

	myList.delete(30)
	myList.print()

}
