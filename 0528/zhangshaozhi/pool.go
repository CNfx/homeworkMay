package main

import (
    "errors"
)

var (
    ErrPoolFull = errors.New("pool full.")
    ErrPoolEmpty = errors.New("pool empty.")
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

func (p *Pool) Pop(x int) error {
    if p.head >= len(p.data) {
       return ErrPoolFull 
    }
    p.data[p.head] = x
    if p.head < len(p.data) - 1 {
        p.head++ 
    }
    return nil
}

func (p *Pool) Push() error {
    if p.head == 0 {
        return ErrPoolEmpty
    }
    p.data[p.head] = 0
    p.head--
    return nil
}
