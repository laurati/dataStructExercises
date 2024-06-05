package main

import "fmt"

type BstNode struct {
	data  int
	left  *BstNode
	right *BstNode
}

func Insert(node *BstNode, value int) *BstNode {
	if node == nil {
		newNode := new(BstNode)
		newNode.data = value
		newNode.left = nil
		newNode.right = nil
		return newNode
	} else if value < node.data {
		node.left = Insert(node.left, value)
	} else if value > node.data {
		node.right = Insert(node.right, value)
	}
	return node
}

func Search(node *BstNode, value int) bool {
	if node == nil {
		return false
	} else if value == node.data {
		return true
	} else if value < node.data {
		return Search(node.left, value)
	} else {
		return Search(node.right, value)
	}
}

func FindMin(node *BstNode) int {
	if node == nil {
		return -1
	} else if node.left == nil {
		return node.data
	}
	return FindMin(node.left)
}

func FindMax(node *BstNode) int {
	if node == nil {
		return -1
	} else if node.right == nil {
		return node.data
	}
	return FindMax(node.right)
}

func main() {

	var tree *BstNode = nil

	tree = Insert(tree, 2)
	tree = Insert(tree, 5)
	tree = Insert(tree, 3)
	tree = Insert(tree, 7)
	tree = Insert(tree, 1)

	fmt.Println(Search(tree, 7))

	fmt.Println(FindMin(tree))

	fmt.Println(FindMax(tree))
}
