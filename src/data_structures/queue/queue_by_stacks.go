// this is a queue implementation with the help of two stacks
package queue

import (
	"errors"
	"golang-dsa/src/data_structures/stack"
)

type QueueByStacks struct {
	s1 *stack.Stack
	s2 *stack.Stack
}

func NewQueueByStacks() *QueueByStacks {
	q := QueueByStacks{}
	q.s1 = stack.New()
	q.s2 = stack.New()

	return &q
}

func (q *QueueByStacks) Push(item interface{}) error {
	return q.s1.Push(item)
}

func (q *QueueByStacks) Pop() (interface{}, error) {
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
