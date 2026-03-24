package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestTree(input []int) *Tree[int] {
	sumFn := func(a ...int) int {
		sum := 0
		for _, x := range a {
			sum += x
		}
		return sum
	}

	compareFn := func(a, b int) int {
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	}

	return NewBinaryTreeWithSlice[int](input, 0, sumFn, compareFn)
}

func TestTraversals(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}

	tests := []struct {
		name     string
		traverse func(*Tree[int], func(*Node[int]))
		expected []int
	}{
		{
			name: "PreOrder",
			traverse: func(t *Tree[int], f func(*Node[int])) {
				t.PreOrderTraversal(f)
			},
			expected: []int{34, 56, 45, 197, 887, 123, 89, 434, 67868},
		},
		{
			name: "PostOrder",
			traverse: func(t *Tree[int], f func(*Node[int])) {
				t.PostOrderTraversal(f)
			},
			expected: []int{197, 887, 45, 123, 56, 434, 67868, 89, 34},
		},
		{
			name: "InOrder",
			traverse: func(t *Tree[int], f func(*Node[int])) {
				t.InOrderTraversal(f)
			},
			expected: []int{197, 45, 887, 56, 123, 34, 434, 89, 67868},
		},
		{
			name: "LevelOrder",
			traverse: func(t *Tree[int], f func(*Node[int])) {
				t.LevelOrderTraversal(f)
			},
			expected: []int{34, 56, 89, 45, 123, 434, 67868, 197, 887},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := newTestTree(input)

			var output []int
			tt.traverse(tree, func(n *Node[int]) {
				output = append(output, n.data)
			})

			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestSize(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := newTestTree(input)

	tree.root.Print()

	assert.Equal(t, len(input), tree.Size())
}

func TestMaxDepth(t *testing.T) {
	input := []int{34, 56, 89, 45, 123, 434, 67868, 197, 887}
	tree := newTestTree(input)

	tree.root.Print()

	assert.Equal(t, tree.MaxDepth(), 3)
}

func TestSumTree(t *testing.T) {
	input := []int{26, 10, 3, 4, 6, 3}
	tree := newTestTree(input)

	tree.root.Print()

	assert.Equal(t, true, tree.IsSumTree())
}
