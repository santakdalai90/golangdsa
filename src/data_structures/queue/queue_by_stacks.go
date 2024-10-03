// this is a queue implementation with the help of two stacks
package queue

import (
	"errors"
	"github.com/santakdalai90/golang-dsa/src/data_structures/stack"
)

type QueueByStacks[T any] struct {
	s1 *stack.Stack[T]
	s2 *stack.Stack[T]
}

func NewQueueByStacks[T any]() *QueueByStacks[T] {
	q := QueueByStacks[T]{}
	q.s1 = stack.New[T]()
	q.s2 = stack.New[T]()

	return &q
}

func (q *QueueByStacks[T]) Push(item T) {
	q.s1.Push(item)
}

func (q *QueueByStacks[T]) Pop() (T, error) {
	front, err := q.s2.Pop()
	if errors.Is(err, stack.NewStackEmptyError()) {
		//s2 is empty; transfer data from s1
		s1f, e := q.s1.Pop()
		for e == nil {
			q.s2.Push(s1f)
			s1f, e = q.s1.Pop()
		}
		front, err = q.s2.Pop()
	}
	return front, err
}
