package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
	sp   int
}

type Queue struct {
	data  []int
	read  int
	write int
}

type LinkNode struct {
	data int
	next int
}

type LinkList struct {
	list []LinkNode
}

var (
	ErrStackFull  = errors.New("stack full")
	ErrStackEmpty = errors.New("stack empty")

	ErrQueueFull  = errors.New("queue full")
	ErrQueueEmpty = errors.New("queue empty")

	ErrLinkListFull  = errors.New("linklist full")
	ErrLinkListEmpty = errors.New("linklist empty")
	ErrLinkListIndex = errors.New("linklist operation index wrong")
)

func NewStack(cap int) *Stack {
	s := make([]int, cap)
	return &Stack{s, 0}
}

func (s *Stack) Pop() (int, error) {
	if s.sp == 0 {
		return 0, ErrStackEmpty
	}

	v := s.data[s.sp-1]
	s.sp--

	return v, nil
}

func (s *Stack) Push(v int) error {
	if s.sp == cap(s.data) {
		return ErrStackFull
	}

	s.data[s.sp] = v
	s.sp++

	return nil
}

func NewQueue(cap int) *Queue {
	q := make([]int, cap)
	return &Queue{q, 0, 0}
}

func (q *Queue) EnQueue(v int) error {
	dataLen := q.write - q.read
	if cap(q.data)-dataLen == 0 {
		return ErrQueueFull
	}

	n := q.write % cap(q.data)
	q.data[n] = v

	q.write++

	return nil
}

func (q *Queue) DeQueue() (int, error) {
	if q.write-q.read == 0 {
		return 0, ErrQueueEmpty
	}

	n := q.read % cap(q.data)
	v := q.data[n]

	q.read++

	return v, nil
}

func NewLinkList(cap int) *LinkList {
	l := make([]LinkNode, cap)

	for i := 0; i < cap; i++ {
		l[i].next = i + 1
	}
	l[cap-1].next = 0

	return &LinkList{l}
}

func LengthList(l *LinkList) int {
	cnt := 0
	i := l.list[cap(l.list)-1].next
	for i > 0 {
		i = l.list[i].next
		cnt++
	}

	return cnt
}

func (l *LinkList) Insert(i, v int) error {
	if i > LengthList(l)+1 {
		return ErrLinkListIndex
	}

	last := cap(l.list) - 1
	freeIndex := l.list[0].next
	if freeIndex > 0 {
		l.list[0].next = l.list[freeIndex].next

		l.list[freeIndex].data = v
		for j := 1; j < i; j++ {
			last = l.list[last].next
		}

		l.list[freeIndex].next = l.list[last].next
		l.list[last].next = freeIndex
	}

	return nil
}

func (l *LinkList) Delete(i int) error {
	if i > LengthList(l)+1 {
		return ErrLinkListIndex
	}

	last := cap(l.list) - 1
	for j := 1; j < i; j++ {
		last = l.list[last].next
	}

	index := l.list[last].next
	l.list[index].data = 0
	l.list[last].next = l.list[index].next

	l.list[index].next = l.list[0].next
	l.list[0].next = index

	return nil
}

//Pool这里没太想明白,使用栈实现了一下思路
type Pool struct {
	data Stack
}

func NewPool(cap int) *Pool {
	s := NewStack(cap)
	for i := 0; i < cap; i++ {
		s.Push(1)
	}
	return &Pool{*s}
}

func (p *Pool) Pop() int {
	v, err := p.data.Pop()
	if err != nil {
		return 0
	}

	return v
}

func (p *Pool) Push(v int) error {
	return p.data.Push(v)
}

func main() {
	s := NewStack(3)

	fmt.Println(s.Push(1), s)
	fmt.Println(s.Push(2), s)
	fmt.Println(s.Push(3), s)
	fmt.Println(s.Push(4), s)

	fmt.Println(s.Pop())
	fmt.Println(s.Push(5), s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	q := NewQueue(3)

	fmt.Println(q.EnQueue(1), q)
	fmt.Println(q.EnQueue(2), q)
	fmt.Println(q.EnQueue(3), q)
	fmt.Println(q.EnQueue(4), q)
	fmt.Println(q.DeQueue())
	fmt.Println(q.EnQueue(5), q)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())

	l := NewLinkList(5)

	fmt.Println(l)
	l.Insert(1, 5)
	fmt.Println(l)
	l.Insert(2, 6)
	fmt.Println(l)
	l.Delete(2)
	fmt.Println(l)
}
