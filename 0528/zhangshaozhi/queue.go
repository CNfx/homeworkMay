package main

import (
    "errors"
)

var (
    ErrQueueFull = errors.New("queue full.")
    ErrQueueEmpty = errors.New("queue empty.")
)

type RingQueue struct {
    data []int
    head int
    tail int
}

func NewQueue(cap int) *RingQueue {
    return &RingQueue{
         data: make([]int, cap),
    }
}

func (q *RingQueue) Push(x int) error {
    if (cap(q.data) - (q.tail - q.head)) == 0 {
        return ErrQueueFull
    }
    n := q.tail % cap(q.data)
    q.data[n] = x
    q.tail++
    return nil
}


func (q *RingQueue) Pop() (int, error) {
    if q.tail == q.head {
        return 0, ErrQueueEmpty
    }
    n := q.head % cap(q.data)
    x := q.data[n]
    q.head++
    return x, nil
}
