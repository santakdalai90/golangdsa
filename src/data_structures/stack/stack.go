package stack

import (
	"fmt"
)

type Stack struct {
	Data []interface{} // top towards the end
}

func New() *Stack {
	s := Stack{}

	s.Data = make([]interface{}, 0)

	return &s
}

func (s *Stack) Push(item interface{}) error {
	s.Data = append(s.Data, item)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s == nil || len(s.Data) == 0 {
		return nil, NewStackEmptyError()
	}
	top := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return top, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s == nil || len(s.Data) == 0 {
		return nil, NewStackEmptyError()
	}
	top := s.Data[len(s.Data)-1]
	return top, nil
}

func (s *Stack) Print() {
	if s == nil || len(s.Data) == 0 {
		fmt.Println("stack is empty")
		return
	}

	for i := len(s.Data) - 1; i >= 0; i-- {
		fmt.Println(s.Data[i])
	}
}
