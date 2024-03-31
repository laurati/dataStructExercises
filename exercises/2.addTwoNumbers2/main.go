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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var dummyHead = &ListNode{}
	current := dummyHead

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10, Next: nil}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry, Next: nil}
	}

	return dummyHead.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Creating the first list: 2 -> 4 -> 3
	var l1 *ListNode
	l1 = appendNode(l1, 2)
	l1 = appendNode(l1, 4)
	l1 = appendNode(l1, 3)

	fmt.Println(l1)

	// Creating the second list: 5 -> 6 -> 4
	var l2 *ListNode
	l2 = appendNode(l2, 5)
	l2 = appendNode(l2, 6)
	l2 = appendNode(l2, 4)

	// Adding the two numbers
	sum := addTwoNumbers(l1, l2)

	// Printing the result: 7 -> 0 -> 8
	printList(sum)

	fmt.Println(sum)
}
