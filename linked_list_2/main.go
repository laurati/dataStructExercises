package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func appendNode(head *ListNode, data int) *ListNode {
	newNode := &ListNode{Val: data, Next: nil}
	if head == nil {
		return newNode
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

func appendNodeInHead(node *ListNode, data int) *ListNode {
	newNode := &ListNode{Val: data, Next: nil}
	if node == nil {
		return newNode
	}
	head := newNode
	head.Next = node

	return head
}

func insertInOrder(head *ListNode, data int) *ListNode {
	newNode := &ListNode{Val: data, Next: nil}

	if head == nil || data <= head.Val {
		newNode.Next = head
		return newNode
	}

	current := head
	for current.Next != nil && current.Next.Val < data {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func searchList(l *ListNode, v int) bool {
	for l != nil {
		if l.Val == v {
			return true
		}
		l = l.Next
	}
	return false

}

func deleteValueList(l *ListNode, v int) {
	if l.Val == v {
		l.Val = l.Next.Val
		l.Next = l.Next.Next
		return
	}
	for l != nil {
		if l.Next.Val == v {
			l.Next = l.Next.Next
			return
		}
		l = l.Next
	}

}

func main() {
	// Creating the first list: 2 -> 3 -> 4
	var l1 *ListNode
	l1 = insertInOrder(l1, 2)
	l1 = insertInOrder(l1, 4)
	l1 = insertInOrder(l1, 3)

	printList(l1)

	// Creating the second list: 5 -> 6 -> 4
	var l2 *ListNode
	l2 = appendNode(l2, 5)
	l2 = appendNode(l2, 6)
	l2 = appendNode(l2, 4)

	printList(l2)

	fmt.Println(searchList(l2, 6))
	fmt.Println(searchList(l2, 7))

	var l3 *ListNode
	l3 = appendNodeInHead(l3, 10)
	l3 = appendNodeInHead(l3, 30)
	l3 = appendNodeInHead(l3, 20)
	l3 = appendNodeInHead(l3, 40)

	printList(l3)

	deleteValueList(l3, 40)

	printList(l3)

}
