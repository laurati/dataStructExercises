package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {

	var tree *BstNode = nil

	tree = Insert(tree, 2)
	tree = Insert(tree, 5)
	tree = Insert(tree, 3)
	tree = Insert(tree, 1)
	tree = Insert(tree, 7)

	assert.NotNil(t, tree)
	assert.Equal(t, 2, tree.data)
	assert.Equal(t, 5, tree.right.data)
	assert.Equal(t, 1, tree.left.data)

	assert.True(t, Search(tree, 5))
	assert.False(t, Search(tree, 20))

}
