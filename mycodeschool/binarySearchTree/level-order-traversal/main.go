package main

import "fmt"

type Node struct {
	data  rune
	left  *Node
	right *Node
}

type Queue struct {
	nodes []*Node
}

func (q *Queue) Enqueue(n *Node) {
	q.nodes = append(q.nodes, n)
}

// Dequeue removes and returns the front element of the queue.
func (q *Queue) Dequeue() *Node {
	if len(q.nodes) == 0 {
		return nil
	}
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

// LevelOrder prints nodes in a binary tree in level order.
func LevelOrder(root *Node) {
	if root == nil {
		return
	}
	queue := Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		fmt.Printf("%c ", current.data)
		if current.left != nil {
			queue.Enqueue(current.left)
		}
		if current.right != nil {
			queue.Enqueue(current.right)
		}
	}
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
	LevelOrder(root)

}
