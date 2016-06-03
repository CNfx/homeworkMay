package stack

import "errors"

type Stack struct {
	bucket []string
	space  int
}

func (s *Stack) Push(data string) error {
	if s.space == 0 {
		return errors.New("Stack full!")
	}

	s.bucket = append(s.bucket, data)
	s.space--
	return nil
}

func (s *Stack) Pop() (string, error) {

	if len(s.bucket) == 0 {
		return "", errors.New("Empty stack")
	}

	data := s.bucket[len(s.bucket)-1]
	s.bucket = s.bucket[:len(s.bucket)-1]
	s.space++

	return data, nil
}

func NewStack(len int) *Stack {
	return &Stack{
		space: len,
	}
}
