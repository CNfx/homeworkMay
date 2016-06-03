package main

import (
    "errors"
    "fmt"
)

type RingQueue struct {
    data []int
    head int
    tail int
}

var (
    ErrQueueFull = errors.New("queue full")
    ErrQueueEmpty = errors.New("queue empty")
)

func NewRingQueue(cap int) *RingQueue  {
    return &RingQueue{
        data : make([]int, cap),
    }
}

func (q *RingQueue) Push(x int) error {
    if (cap(q.data) - (q.tail - q.head)) == 0{
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

func main(){
    s := NewRingQueue(3)
    
    fmt.Println(s.Push(1), s)
    fmt.Println(s.Push(2), s)
    fmt.Println(s.Push(3), s)
    fmt.Println(s.Push(4), s)
    
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    fmt.Println(s.Push(4),s)
    fmt.Println(s.Push(5),s)
    fmt.Println(s.Pop())
}
