package main

import (
	"fmt"
	"math"
)

const (
	minValue = math.MinInt64
	maxValue = math.MaxInt64
)

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

// func FindMin(root *BstNode) int {
// 	if root == nil {
// 		return -1
// 	}
// 	var current *BstNode = root
// 	for current.left != nil {
// 		current = current.left
// 	}
// 	return current.data
// }

func FindMin(root *BstNode) int {
	if root == nil {
		return -1
	} else if root.left == nil {
		return root.data
	}
	return FindMin(root.left)
}

// func FindMax(root *BstNode) int {
// 	if root == nil {
// 		return -1
// 	}
// 	var current *BstNode = root
// 	for current.right != nil {
// 		current = current.right
// 	}
// 	return current.data
// }

func FindMax(root *BstNode) int {
	if root == nil {
		return -1
	} else if root.right == nil {
		return root.data
	}
	return FindMax(root.right)
}

func FindHeight(root *BstNode) int {
	if root == nil {
		return -1
	}
	left := float64(FindHeight(root.left))
	right := float64(FindHeight(root.right))

	return int(math.Max(left, right)) + 1
}

func IsBst(root *BstNode, min, max int64) bool {
	if root == nil {
		return true
	}
	if int64(root.data) > min &&
		int64(root.data) < max &&
		IsBst(root.left, min, int64(root.data)) &&
		IsBst(root.right, int64(root.data), max) {
		return true
	}
	return false
}

// FindMin encontra o nó com o menor valor na árvore
func FindMinNode(root *BstNode) *BstNode {
	for root.left != nil {
		root = root.left
	}
	return root
}

// Delete exclui um nó da árvore com o valor fornecido
func Delete(root *BstNode, data int) *BstNode {
	if root == nil {
		return root
	} else if data < root.data {
		root.left = Delete(root.left, data)
	} else if data > root.data {
		root.right = Delete(root.right, data)
	} else {
		// Caso 1: Nenhum filho
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil { // Caso 2: Um filho
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else { // Caso 3: Dois filhos
			temp := FindMinNode(root.right)
			root.data = temp.data
			root.right = Delete(root.right, temp.data)
		}
	}
	return root
}

// Find encontra um nó com o valor fornecido na árvore
func Find(root *BstNode, data int) *BstNode {
	if root == nil {
		return nil
	} else if root.data == data {
		return root
	} else if data < root.data {
		return Find(root.left, data)
	} else {
		return Find(root.right, data)
	}
}

// GetSuccessor encontra o sucessor em ordem de um nó na BST
func GetSuccessor(root *BstNode, data int) *BstNode {
	// Busca o nó - O(h)
	current := Find(root, data)
	if current == nil {
		return nil
	}
	// Caso 1: O nó tem subárvore direita
	if current.right != nil {
		return FindMinNode(current.right) // O(h)
	} else { // Caso 2: O nó não tem subárvore direita - O(h)
		var successor *BstNode
		ancestor := root
		for ancestor != current {
			if current.data < ancestor.data {
				successor = ancestor // até agora este é o nó mais profundo para o qual o nó atual está à esquerda
				ancestor = ancestor.left
			} else {
				ancestor = ancestor.right
			}
		}
		return successor
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

	fmt.Println(Search(root, 20))

	fmt.Println(FindMin(root))
	fmt.Println(FindMax(root))

	fmt.Println(FindHeight(root))

	fmt.Println(IsBst(root, minValue, maxValue))
}
