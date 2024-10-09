package linkedlist

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	data T
	next *Node[T]
}

func NewNode[T constraints.Ordered](data T, next *Node[T]) *Node[T] {
	return &Node[T]{data, next}
}

func (n *Node[T]) GetData() T {
	if n == nil {
		panic("cannot return data of a nil node")
	}
	return n.data
}

func (n *Node[T]) GetNext() *Node[T] {
	if n != nil {
		return n.next
	}
	return nil
}

func (n *Node[T]) SetData(data T) {
	if n == nil {
		panic("cannot set data of a nil node")
	}
	n.data = data
}

func (n *Node[T]) SetNext(next *Node[T]) {
	if n == nil {
		panic("cannot set next of a nil node")
	}
	n.next = next
}
