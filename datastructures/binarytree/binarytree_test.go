package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrderTraversal(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](
		input,
		0,
		func(a ...int) int {
			sum := 0
			for _, x := range a {
				sum += x
			}
			return sum
		},
		func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	)

	tree.root.Print()
	output := make([]int, 0)
	tf := func(n *Node[int]) {
		output = append(output, n.data)
	}
	tree.PreOrderTraversal(tf)

	assert.Equal(t, len(output), 9)
	t.Log(output)
	assert.Equal(t, []int{34, 56, 45, 197, 887, 123, 89, 434, 67868}, output)
}

func TestPostOrderTraversal(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](
		input,
		0,
		func(a ...int) int {
			sum := 0
			for _, x := range a {
				sum += x
			}
			return sum
		},
		func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	)

	tree.root.Print()
	output := make([]int, 0)
	tf := func(n *Node[int]) {
		output = append(output, n.data)
	}
	tree.PostOrderTraversal(tf)

	assert.Equal(t, len(output), 9)
	t.Log(output)
	assert.Equal(t, []int{197, 887, 45, 123, 56, 434, 67868, 89, 34}, output)
}

func TestInOrderTraversal(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](
		input,
		0,
		func(a ...int) int {
			sum := 0
			for _, x := range a {
				sum += x
			}
			return sum
		},
		func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	)

	tree.root.Print()
	output := make([]int, 0)
	tf := func(n *Node[int]) {
		output = append(output, n.data)
	}
	tree.InOrderTraversal(tf)

	assert.Equal(t, len(output), 9)
	t.Log(output)
	assert.Equal(t, []int{197, 45, 887, 56, 123, 34, 434, 89, 67868}, output)
}

func TestLevelOrderTraversal(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](
		input,
		0,
		func(a ...int) int {
			sum := 0
			for _, x := range a {
				sum += x
			}
			return sum
		},
		func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	)

	tree.root.Print()
	output := make([]int, 0)
	tf := func(n *Node[int]) {
		output = append(output, n.data)
	}
	tree.LevelOrderTraversal(tf)

	assert.Equal(t, len(output), 9)
	t.Log(output)
	assert.Equal(t, []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}, output)
}

func TestSize(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](
		input,
		0,
		func(a ...int) int {
			sum := 0
			for _, x := range a {
				sum += x
			}
			return sum
		},
		func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	)

	tree.root.Print()
	size := tree.Size()

	assert.Equal(t, size, len(input))
}

func TestMaxDepth(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := NewBinaryTreeWithSlice[int](
		input,
		0,
		func(a ...int) int {
			sum := 0
			for _, x := range a {
				sum += x
			}
			return sum
		},
		func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	)

	tree.root.Print()
	depth := tree.MaxDepth()

	assert.Equal(t, depth, 3)
}

func TestSumTree(t *testing.T) {
	input := []int{26, 10, 3, 4, 6, 3}
	tree := NewBinaryTreeWithSlice[int](
		input,
		0,
		func(a ...int) int {
			sum := 0
			for _, x := range a {
				sum += x
			}
			return sum
		},
		func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		},
	)

	tree.root.Print()
	output := tree.IsSumTree()

	assert.Equal(t, true, output)
}
