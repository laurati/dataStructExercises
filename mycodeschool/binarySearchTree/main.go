package main

import "fmt"

type BstNode struct {
	data  int
	left  *BstNode
	right *BstNode
}

func GetNewNode(data int) *BstNode {
	newNode := new(BstNode)
	newNode.data = data
	newNode.left = nil
	newNode.right = nil
	return newNode
}

func Insert(root *BstNode, data int) *BstNode {
	if root == nil {
		root = GetNewNode(data)
		return root
	} else if data <= root.data {
		root.left = Insert(root.left, data)
	} else {
		root.right = Insert(root.right, data)
	}
	return root
}

func Search(root *BstNode, data int) bool {
	if root == nil {
		return false
	} else if root.data == data {
		return true
	} else if data <= root.data {
		return Search(root.left, data)
	} else {
		return Search(root.right, data)
	}

}

func main() {

	var root *BstNode = nil

	root = Insert(root, 15)
	root = Insert(root, 10)
	root = Insert(root, 20)
	root = Insert(root, 25)
	root = Insert(root, 8)
	root = Insert(root, 12)

	fmt.Println(root)
	// fmt.Println("Left Child of Root:", root.left, &root.left.data)
	// fmt.Println("Right Child of Root:", root.right, &root.right.data)

	fmt.Println(Search(root, 20))
}
