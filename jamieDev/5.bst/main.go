package main

import "fmt"

var count int

//Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert will add a node to the tree
//the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}

}

//The plan is to compare the key with the receiver and go to the child
//on the right if it's lager and go to the child on the left if
//it is smaller

//====================================================

//Search will take in a key value
//and return true if there is a node with that value
func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)

	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	// 		 100
	// 		/   \
	//    52    203
	//   / \     / \
	//  19 76   150 310
	//  / \   \     /
	//  7  24  88  276

	fmt.Println(tree.Search(276))
	fmt.Println(count)

}
