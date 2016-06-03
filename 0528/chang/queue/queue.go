package queue

import (
	"errors"
	"sync"
)

var (
	ErrQueueFull  = errors.New("queue full")
	ErrQueueEmpty = errors.New("queue empty")
)

type Queue struct {
	data []int
	head int
	tail int
	lock sync.Mutex
}

func NewQueue(cap int) *Queue {
	return &Queue{
		data: make([]int, cap),
	}
}

func (q *Queue) Push(v int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if (cap(q.data) - (q.tail - q.head)) == 0 {
		return ErrQueueFull
	}

	n := q.tail % cap(q.data)
	q.data[n] = v

	q.tail++
	return nil
}

func (q *Queue) Pop() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.tail == q.head {
		return 0, ErrQueueEmpty
	}

	n := q.head % cap(q.data)
	v := q.data[n]

	q.head++
	return v, nil
}
