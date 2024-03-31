package main

import (
	"fmt"
	"strconv"
)

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

func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {

	var l1String string
	toPrint1 := l1.head
	for l1.length != 0 {
		l1String += strconv.Itoa(toPrint1.data)
		toPrint1 = toPrint1.next
		l1.length--
	}

	var l2String string
	toPrint2 := l2.head
	for l2.length != 0 {
		l2String += strconv.Itoa(toPrint2.data)
		toPrint2 = toPrint2.next
		l2.length--
	}

	l1int, _ := strconv.Atoi(l1String)
	l2int, _ := strconv.Atoi(l2String)

	sum := l1int + l2int
	newSum := strconv.Itoa(sum)

	newList := LinkedList{}

	for i := 0; i < len(newSum); i++ {
		digito, _ := strconv.Atoi(string(newSum[i]))
		node := node{data: digito}
		newList.prepend(&node)
	}

	return &newList
}

func main() {
	node1 := node{data: 2}
	node2 := node{data: 4}
	node3 := node{data: 3}

	myList1 := LinkedList{}

	myList1.prepend(&node1)
	myList1.prepend(&node2)
	myList1.prepend(&node3)

	// myList1.print()

	// fmt.Println("==========")

	node4 := node{data: 5}
	node5 := node{data: 6}
	node6 := node{data: 4}

	myList2 := LinkedList{}

	myList2.prepend(&node4)
	myList2.prepend(&node5)
	myList2.prepend(&node6)

	// myList2.print()

	newList := addTwoNumbers(&myList1, &myList2)
	newList.print()

}
