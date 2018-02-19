package stack

import (
	"errors"
	ll "github.com/eyoung8/go-fun/linkedlist"
)

var (
	EmptyStackErr = errors.New("Empty stack")
)

type Stack struct {
	list ll.LinkedList
}

func (s *Stack) Pop() (interface{}, error) {
	val := s.list.Pop()
	if val == nil {
		return nil, EmptyStackErr
	}
	return val, nil
}

func (s *Stack) Push(val interface{}) {
	s.list.AddBack(val)
}

func (s *Stack) Peek() (interface{}, error) {
	val := s.list.Back()
	if val == nil {
		return nil, EmptyStackErr
	}
	return val, nil
}

func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}
