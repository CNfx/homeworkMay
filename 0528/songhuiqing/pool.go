package main

import (
    "errors"
)

var (
    ErrPoolFull = errors.New("pool full")
    ErrPoolEmpty = errors.New("pool empty")
)

type Pool struct {
    head int
    data []int
}

func NewPool(cap int) *Pool {
    return &Pool{
        data: make([]int, cap),
    }
}

func (p *Pool) Push(x int) error {
    if p.head >= len(p.data) {
       return ErrPoolFull 
    }
    p.data[p.head] = x
    if p.head < len(p.data) - 1 {
        p.head++ 
    }
    return nil
}

func (p *Pool) Pop() (int, error) {
    if p.head == 0 {
        return 0, ErrPoolEmpty
    }
    i := p.head - 1
    d := p.data[i]
    p.head = i
    //println(d)    
    return d, nil
}


func main() {
    p := NewPool(10)
    
    p.Push(100)
    p.Push(200)
    p.Push(300)
    
    p.Pop()
    p.Pop()
    p.Pop()
    p.Push(999)
    p.Pop()
}