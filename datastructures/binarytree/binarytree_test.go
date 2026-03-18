package binarytree

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomBinaryTree(t *testing.T) {
	nodeCount := 25
	NewRandomBinaryTree[int](nodeCount, func() int {
		return rand.Intn(100)
	})
}

func TestPreOrderTraversal(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](input, 0)

	tree.Print()
	output := make([]int, 0)
	tf := func(n *Node[int]) {
		output = append(output, n.data)
	}
	tree.PreOrderTraversal(tf)

	assert.Equal(t, len(output), 9)
	t.Log(output)
	assert.Equal(t, []int{34, 56, 45, 197, 887, 123, 89, 434, 67868}, output)
}

func TestInOrderTraversal(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](input, 0)

	tree.Print()
	output := make([]int, 0)
	tf := func(n *Node[int]) {
		output = append(output, n.data)
	}
	tree.InOrderTraversal(tf)

	assert.Equal(t, len(output), 9)
	t.Log(output)
	assert.Equal(t, []int{197, 45, 887, 56, 123, 34, 434, 89, 67868}, output)
}

func TestPostOrderTraversal(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](input, 0)

	tree.Print()
	output := make([]int, 0)
	tf := func(n *Node[int]) {
		output = append(output, n.data)
	}
	tree.PostOrderTraversal(tf)

	assert.Equal(t, len(output), 9)
	t.Log(output)
	assert.Equal(t, []int{197, 887, 45, 123, 56, 434, 67868, 89, 34}, output)
}
