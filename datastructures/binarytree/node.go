package binarytree

import (
	"fmt"
	"math/rand"
)

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

func (b *Node[T]) Delete() *Node[T] {
	if b == nil {
		return nil
	}

	b.left = b.left.Delete()
	b.right = b.right.Delete()
	return nil
}
