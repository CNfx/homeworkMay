package main

import "sync"

type Stack struct {
	stack []int
	len   int
	lock  sync.Mutex
}

func New() *Stack {
	s := &Stack{}
	s.stack = make([]int, 0)
	s.len = 0

	return s
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	return s.len
}

func (s *Stack) isEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	return s.len == 0
}

func (s *Stack) Push(el int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	prepend := make([]int, 1)
	prepend[0] = el
	s.stack = append(prepend, s.stack...)
	s.len++
}

func (s *Stack) Pop() (el int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	el, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}

func (s *Stack) Peek() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	return s.stack[0]
}

func main() {
	s := New()	
	println(s.isEmpty())
	s.Push(100)
	s.Push(101)
	s.Push(102)
	s.Push(103)
	s.Push(104)
	println(s.isEmpty())
	println(s.Len())
	println(s.Pop())
	println(s.Len())
	println(s.Peek())
}
