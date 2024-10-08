package linkedlist

import "fmt"

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
	newNode := NewNode(arr[0], nil) // &Node[T]{data: arr[0]}
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
		// curr.SetNext(&Node[T]{data: a})
		curr.SetNext(NewNode(a, nil))
		curr = curr.GetNext()
	}
}

func (sll *SinglyLinkedList[T]) GetHead() *Node[T] {
	return sll.head
}

func (sll *SinglyLinkedList[T]) AppendToTail(data T) {
	newNode := NewNode(data, nil) //&Node[T]{data: data}

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
