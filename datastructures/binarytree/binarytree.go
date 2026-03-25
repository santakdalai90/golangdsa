package binarytree

import (
	"log"
)

type Tree[T any] struct {
	root    *Node[T]
	zero    T
	add     func(...T) T
	compare func(T, T) int // -1 if less, 0 if equal, 1 if more
}

func New[T any](rootValue, zero T, add func(...T) T, compare func(T, T) int) *Tree[T] {
	return &Tree[T]{
		root:    &Node[T]{data: rootValue},
		zero:    zero,
		add:     add,
		compare: compare,
	}
}

func NewRandomBinaryTree[T any](n int, zero T, add func(...T) T, compare func(T, T) int, generator func() T) *Tree[T] {
	log.Printf("generating random tree with %d nodes", n)
	var root = NewNode(generator())
	log.Printf("created root node with data %v", root.data)
	for range n - 1 {
		root.Insert(NewNode(generator()))
	}
	log.Printf("generated random tree with %d nodes", n)

	return &Tree[T]{
		root:    root,
		zero:    zero,
		add:     add,
		compare: compare,
	}
}

func NewBinaryTreeWithSlice[T any](data []T, zero T, add func(...T) T, compare func(T, T) int) *Tree[T] {
	if len(data) == 0 {
		return nil
	}

	queue := make([]*Node[T], 1, len(data))
	queue[0] = NewNode(data[0])

	front := 0
	i := 1

	for front < len(queue) && i < len(data) {
		var node = queue[front]
		front++

		node.left = NewNode(data[i])
		queue = append(queue, node.left)
		i++

		if i < len(data) {
			node.right = NewNode(data[i])
			queue = append(queue, node.right)
			i++
		}
	}

	return &Tree[T]{
		root:    queue[0],
		zero:    zero,
		add:     add,
		compare: compare,
	}
}

func (t *Tree[T]) PreOrderTraversal(tf TraverseFunc[T]) {
	t.root.PreOrderTraversal(tf)
}

func (t *Tree[T]) PostOrderTraversal(tf TraverseFunc[T]) {
	t.root.PostOrderTraversal(tf)
}

func (t *Tree[T]) InOrderTraversal(tf TraverseFunc[T]) {
	t.root.InOrderTraversal(tf)
}

func (t *Tree[T]) MorrisInOrderTraversal(tf TraverseFunc[T]) {
	t.root.MorrisInOrderTraversal(tf)
}

func (t *Tree[T]) LevelOrderTraversal(tf TraverseFunc[T]) {
	t.root.LevelOrderTraversal(tf)
}

func (t *Tree[T]) MaxDepth() int {
	return t.root.MaxDepth()
}

func (t *Tree[T]) Size() int {
	return t.root.Size()
}

func (t *Tree[T]) Max() T {
	if maxNode := t.root.Max(t.compare); maxNode != nil {
		return maxNode.data
	}
	return t.zero
}

func (t *Tree[T]) IsSumTree() bool {
	return t.root.IsSumTree(t.add, t.zero, t.compare)
}
