package main

import "fmt"

type Node struct {
	data  rune
	left  *Node
	right *Node
}

func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%c ", root.data)
	Preorder(root.left)
	Preorder(root.right)
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.left)
	fmt.Printf("%c ", root.data)
	Inorder(root.right)
}

func Postorder(root *Node) {
	if root == nil {
		return
	}
	Postorder(root.left)
	Postorder(root.right)
	fmt.Printf("%c ", root.data)

}

// Insert inserts a node into the binary search tree.
func Insert(root *Node, data rune) *Node {
	if root == nil {
		root = &Node{data: data}
	} else if data <= root.data {
		root.left = Insert(root.left, data)
	} else {
		root.right = Insert(root.right, data)
	}
	return root
}

func main() {

	// Code to test the logic
	// Creating an example tree
	//         M
	//       /   \
	//      B     Q
	//     / \     \
	//    A   C     Z

	var root *Node
	root = Insert(root, 'M')
	root = Insert(root, 'B')
	root = Insert(root, 'Q')
	root = Insert(root, 'Z')
	root = Insert(root, 'C')
	root = Insert(root, 'A')

	// Print nodes in level order
	Preorder(root)

}
