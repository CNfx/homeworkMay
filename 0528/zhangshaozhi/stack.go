package main

import (
    "errors"
)

var (
    ErrStackFull = errors.New("stack full")
    ErrStackEmpty = errors.New("stack empty")
)

type Stack struct {
    ptr []int
    sp int
}

func NewStack(cap, sp int) *Stack {
   return &Stack{ptr: make([]int, cap), sp: sp}
}

func (s *Stack) Push(v int) error {
    if s.sp == 0 {
        return ErrStackFull
    }
    s.sp--
    s.ptr[s.sp] = v
    return nil
}

func (s *Stack) Pop() (int, error) {
    if s.sp == len(s.ptr) {
        return 0, ErrStackEmpty
    }
    v := s.ptr[s.sp]
    s.sp++
    return v, nil
}
