package binarytree

import (
	"fmt"
	"log"
	"math/rand"
)

type TraverseFunc[T any] func(node *Node[T])

type Node[T any] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		data: data,
	}
}

func (b *Node[T]) PreOrderTraversal(tf TraverseFunc[T]) {
	if b == nil {
		return
	}

	tf(b)
	b.left.PreOrderTraversal(tf)
	b.right.PreOrderTraversal(tf)
}

func (b *Node[T]) PostOrderTraversal(tf TraverseFunc[T]) {
	if b == nil {
		return
	}

	b.left.PostOrderTraversal(tf)
	b.right.PostOrderTraversal(tf)
	tf(b)
}

func (b *Node[T]) InOrderTraversal(tf TraverseFunc[T]) {
	if b == nil {
		return
	}

	b.left.InOrderTraversal(tf)
	tf(b)
	b.right.InOrderTraversal(tf)
}

func (b *Node[T]) Insert(node *Node[T]) {
	if b == nil || node == nil {
		return
	}

	shouldGoLeft := rand.Intn(2) == 1
	if shouldGoLeft {
		if b.left != nil {
			b.left.Insert(node)
		} else {
			b.left = node
		}
	} else {
		if b.right != nil {
			b.right.Insert(node)
		} else {
			b.right = node
		}
	}
}

func (b *Node[T]) InsertLeft(node *Node[T]) (*Node[T], error) {
	if b == nil {
		return nil, fmt.Errorf("target node is nil")
	}

	if node == nil {
		return nil, fmt.Errorf("input node is nil")
	}

	if b.left != nil {
		return nil, fmt.Errorf("left of target node is already occupied")
	}
	b.left = node
	return node, nil
}

func (b *Node[T]) InsertRight(node *Node[T]) (*Node[T], error) {
	if b == nil {
		return nil, fmt.Errorf("target node is nil")
	}

	if node == nil {
		return nil, fmt.Errorf("input node is nil")
	}

	if b.right != nil {
		return nil, fmt.Errorf("right of target node is already occupied")
	}
	b.right = node
	return node, nil
}

func NewRandomBinaryTree[T any](n int, generator func() T) *Node[T] {
	log.Printf("generating random tree with %d nodes", n)
	var root = NewNode(generator())
	log.Printf("created root node with data %v", root.data)
	for range n - 1 {
		root.Insert(NewNode(generator()))
	}
	log.Printf("generated random tree with %d nodes", n)

	return root
}

func NewBinaryTreeWithSlice[T any](data []T, index int) *Node[T] {
	var root *Node[T]

	if index < len(data) {
		root = &Node[T]{data: data[index]}
		root.left = NewBinaryTreeWithSlice(data, 2*index+1)
		root.right = NewBinaryTreeWithSlice(data, 2*index+2)
	}
	return root
}
