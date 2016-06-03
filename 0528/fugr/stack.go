package structure

import (
	"errors"
)

var (
	ErrStackFull  = errors.New("stack full")
	ErrStackEmpty = errors.New("stack empty")
)

type Stack []int

func NewStack(cap int) Stack {
	s := make([]int, cap+1)
	s[0] = cap
	return s
}

func (s Stack) Push(dat int) error {
	pos := s[0]
	if pos == 0 {
		return ErrStackFull
	}

	s[pos] = dat
	s[0]--

	return nil
}

func (s Stack) Pop() (int, error) {
	pos := s[0] + 1
	if pos == len(s) {
		return 0, ErrStackEmpty
	}

	dat := s[pos]

	s[0]++
	return dat, nil
}
