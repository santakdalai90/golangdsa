package stack

import (
	"fmt"
)

type Stack[T any] struct {
	data []T
	size int
}

func New[T any]() *Stack[T] {
	s := new(Stack[T])
	s.data = make([]T, 0)
	s.size = 0
	return s
}

func (s *Stack[T]) Push(k T) {
	s.data = append(s.data, k)
	s.size++
}

func (s *Stack[T]) Top() (T, error) {
	if s.size == 0 {
		var empty T
		return empty, NewStackEmptyError()
	}
	return s.data[s.size-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size == 0 {
		var empty T
		return empty, NewStackEmptyError()
	}
	p := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return p, nil
}

func (s *Stack[T]) Print() {
	if s == nil || len(s.data) == 0 {
		fmt.Println("stack is empty")
		return
	}

	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}
