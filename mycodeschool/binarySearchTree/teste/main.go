package main

import "fmt"

type BstNode struct {
	data  int
	left  *BstNode
	right *BstNode
}

func (node *BstNode) Insert(value int) {
	if node == nil {
		return
	}
	if value < node.data {
		if node.left == nil {
			node.left = &BstNode{data: value}
		} else {
			node.left.Insert(value)
		}
	} else if value > node.data {
		if node.right == nil {
			node.right = &BstNode{data: value}
		} else {
			node.right.Insert(value)
		}
	}
}

func (node *BstNode) Search(value int) bool {
	if node == nil {
		return false
	} else if value == node.data {
		return true
	} else if value <= node.data {
		return node.left.Search(value)
	} else {
		return node.right.Search(value)
	}
}

// func FindMin(node *BstNode) int {

// }

func main() {

	var tree *BstNode = &BstNode{data: 2}

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)

	fmt.Println(tree.Search(7))

}
