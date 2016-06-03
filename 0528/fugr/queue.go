package structure

import (
	"errors"
	"sync"
)

var (
	ErrQueueFull  = errors.New("queue full")
	ErrQueueEmpty = errors.New("queue empty")
)

type Queue struct {
	head int
	tail int
	data []int
}

func NewQueue(cap int) Queue {
	return Queue{
		data: make([]int, cap),
	}
}

func (q *Queue) Pop() (int, error) {
	if q.head == q.tail {
		return 0, ErrQueueEmpty
	}

	pos := q.tail % len(q.data)
	q.tail++

	return q.data[pos], nil
}

func (q *Queue) Push(dat int) error {
	if q.head-q.tail == len(q.data) {
		return ErrQueueFull
	}

	pos := q.head % len(q.data)
	q.data[pos] = dat
	q.head++

	return nil
}

type SafeQueue struct {
	lock  *sync.Mutex
	queue Queue
}

func NewSafeQueue(cap int) *SafeQueue {
	return &SafeQueue{
		lock: new(sync.Mutex),
		queue: Queue{
			data: make([]int, cap),
		},
	}
}

func (q *SafeQueue) Pop() (int, error) {
	q.lock.Lock()
	dat, err := q.queue.Pop()
	q.lock.Unlock()

	return dat, err
}

func (q *SafeQueue) Push(dat int) error {
	q.lock.Lock()
	err := q.queue.Push(dat)
	q.lock.Unlock()

	return err
}
