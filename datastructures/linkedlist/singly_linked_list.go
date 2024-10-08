package linkedlist

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (sll *SinglyLinkedList[T]) AppendArray(arr []T) {
	if len(arr) == 0 {
		return
	}

	// append first element from the array
	newNode := &Node[T]{data: arr[0]}
	if sll.GetHead() == nil {
		sll.head = newNode
	} else {
		curr := sll.GetHead()
		for curr.GetNext() != nil {
			curr = curr.GetNext()
		}
		curr.SetNext(newNode)
	}

	// append remainder of the array
	curr := newNode
	for _, a := range arr[1:] {
		curr.SetNext(&Node[T]{data: a})
		curr = curr.GetNext()
	}
}

func (sll *SinglyLinkedList[T]) GetHead() *Node[T] {
	return sll.head
}

func (sll *SinglyLinkedList[T]) AppendToTail(data T) {
	newNode := &Node[T]{data: data}

	if sll.head == nil {
		sll.head = newNode
		return
	}

	curr := sll.head
	for curr.GetNext() != nil {
		curr = curr.GetNext()
	}

	curr.SetNext(newNode)
}

func (sll *SinglyLinkedList[T]) DeleteNode(data T) {
	if sll.head == nil {
		return
	}

	if sll.head.data == data {
		sll.head = sll.head.GetNext()
		return
	}

	curr := sll.head
	for curr.GetNext() != nil {
		if curr.GetNext().GetData() == data {
			curr.SetNext(curr.GetNext().GetNext())
			return
		}
		curr = curr.GetNext()
	}
}

func (sll *SinglyLinkedList[T]) Display() {
	curr := sll.head
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.GetNext()
	}
	fmt.Println(nil)
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
