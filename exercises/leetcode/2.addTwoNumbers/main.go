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
		// Se um dos ponteiros estiver nulo, significa que uma lista é mais curta do que a outra,
		// então assumimos o valor do nó como zero.
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
		// Quando o valor de sum é menor que 10, o operador % (módulo) simplesmente retorna o próprio valor de sum,
		// pois não há "resto" ao dividir por 10
		// current é um ponteiro que aponta para dummyHead
		// sum % 10: garantir que seja um único dígito
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

	// Creating the second list: 5 -> 6 -> 4
	var l2 *ListNode
	l2 = appendNode(l2, 5)
	l2 = appendNode(l2, 6)
	l2 = appendNode(l2, 4)

	lala := addTwoNumbers(l1, l2)
	printList(lala)
}
