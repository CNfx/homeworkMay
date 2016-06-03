package main

import "sync"

type Queue struct {
	queue []int
	len   int
	lock  *sync.Mutex
}

func New() *Queue {
	queue := &Queue{}
	queue.queue = make([]int, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)

	return queue
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len == 0
}

func (q *Queue) Shift() (el int) {
	el, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

func (q *Queue) Push(el int) {
	q.queue = append(q.queue, el)
	q.len++

	return
}

func (q *Queue) Peek() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queue[0]
}
