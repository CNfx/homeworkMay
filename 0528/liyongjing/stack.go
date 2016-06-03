package main

import (
    "errors"
    "fmt"
)

type Stack []int

var (
    ErrStackFull = errors.New("stack full")
    ErrStackEmpty = errors.New("stack empty")
)

func NewStack(cap int) Stack {
    s := make([]int, cap + 1)
    s[0] = cap
    return s
}

func (s Stack) Push(v int) error  {
    i := s[0]
    if i==0 {
        return ErrStackFull
    }
    s[i] = v
    s[0] = i-1
    return nil
}

func (s Stack) Pop() (int, error) {
    i :=s[0] +1
    if i==cap(s){
        return 0,ErrStackEmpty
    }
    v :=s[i]
    s[0] = i
    return v, nil
}

func main(){
    s := NewStack(3)
    
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