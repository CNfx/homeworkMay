package stack

import (
	"errors"
	"sync"
)

var (
	ErrStackFull  = errors.New("stack full")
	ErrStackEmpty = errors.New("stack empty")
)

func NewStack(cap int) *Stack {
	return &Stack{data: make([]int, 0, cap)}
}

type Stack struct {
	data []int
	lock sync.Mutex
}

func (s *Stack) Push(n int) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.data) == cap(s.data) {
		return ErrStackFull
	}

	s.data = append(s.data, n)
	return nil
}

func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.data) == 0 {
		return 0, ErrStackEmpty
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}
