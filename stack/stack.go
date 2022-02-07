package stack

import "errors"

var (
	ErrStackOverflow   = errors.New("stack overflow")
	ErrInvalidCapacity = errors.New("capacity must be > 0")
	ErrEmptyStack      = errors.New("stack is empty")
)

type Stack[T any] struct {
	top  int
	data []*T
}

func New[T any](capacity int) (*Stack[T], error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}

	return &Stack[T]{
		top:  -1,
		data: make([]*T, capacity),
	}, nil
}

func (s *Stack[T]) Push(x *T) error {
	if s.top+1 == len(s.data) {
		return ErrStackOverflow
	}

	s.top++
	s.data[s.top] = x
	return nil
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.top == -1 {
		return s.data[0], ErrEmptyStack
	}

	x := s.data[s.top]
	s.top--
	return x, nil
}

func (s *Stack[T]) Peek() (*T, error) {
	if s.top == -1 {
		return s.data[0], ErrEmptyStack
	}

	return s.data[s.top], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack[T]) IsFull() bool {
	return s.top+1 == len(s.data)
}
