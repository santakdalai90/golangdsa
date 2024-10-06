package linkedlist

import "fmt"

type node[T comparable] struct {
	data T
	next *node[T]
}

type singlyLinkedList[T comparable] struct {
	head *node[T]
}

func NewSinglyLinkedList[T comparable]() *singlyLinkedList[T] {
	return &singlyLinkedList[T]{
		&node[T]{},
	}
}

func (sll *singlyLinkedList[T]) GetFirstNode() *node[T] {
	return sll.head.next // skip the dummy head
}

func (sll *singlyLinkedList[T]) AppendToTail(data T) {
	newNode := &node[T]{data: data}

	curr := sll.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (sll *singlyLinkedList[T]) DeleteNode(data T) {
	curr := sll.head
	for curr.next != nil {
		if curr.next.data == data {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func (sll *singlyLinkedList[T]) Display() {
	curr := sll.head.next // skip the dummy head node
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.next
	}
	fmt.Println(nil)
}
