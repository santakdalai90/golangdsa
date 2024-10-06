package linkedlist

import "fmt"

type node[T comparable] struct {
	data T
	next *node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *node[T]
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (sll *SinglyLinkedList[T]) GetHead() *node[T] {
	return sll.head
}

func (sll *SinglyLinkedList[T]) AppendToTail(data T) {
	newNode := &node[T]{data: data}

	if sll.head == nil {
		sll.head = newNode
		return
	}

	curr := sll.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (sll *SinglyLinkedList[T]) DeleteNode(data T) {
	if sll.head == nil {
		return
	}

	if sll.head.data == data {
		sll.head = sll.head.next
		return
	}

	curr := sll.head
	for curr.next != nil {
		if curr.next.data == data {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func (sll *SinglyLinkedList[T]) Display() {
	curr := sll.head
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.next
	}
	fmt.Println(nil)
}

func (n *node[T]) GetData() T {
	if n == nil {
		panic("cannot return data of a nil node")
	}
	return n.data
}

func (n *node[T]) GetNext() *node[T] {
	if n != nil {
		return n.next
	}
	return n.next
}
