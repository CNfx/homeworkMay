package main

import (
    //"fmt"
)

func main() {
    /*
    //Stack:
    s := NewStack(5, 5)
    for i := 0; i < 6; i++ {
        err := s.Push(i)
        if err != nil {
            fmt.Println(err)
        }
    }
    fmt.Println(s)
    for i := 0; i < 10; i++ {
        v, err := s.Pop()
        if err != nil {
            fmt.Println(err)
        }else{
            fmt.Println(v)
        }
    }
    fmt.Println(s)
    */
    
    /*
    // RingQueue:
    s := NewQueue(5)
    for i := 0; i <= 5; i++ {
        if err := s.Push(i); err != nil {
            fmt.Println(err)
        }
    }
    fmt.Println(s)
    for i := 0; i <= 5; i++ {
        if val, err := s.Pop(); err != nil{
            fmt.Println(err)
        } else {
            fmt.Println(val)
        }
    }
    fmt.Println(s)
    */

    /*
    // linked-list:
    var l LinkedList
    for i := 0; i < 3; i++ {
        l.Append(3)
    }
    l.Insert(5, 1)
    l.Insert(8, 4)
    data := l.Search()
    fmt.Println("main:", data)
    l.Remove(2)
    data = l.Search()
    fmt.Println("main:", data)
    */
    
    /*
    // Pool
    pool := NewPool(10)
    for i := 0; i < 10; i++ {
        if ok := pool.Pop(i); ok != nil {
            fmt.Println(ok)
        }
    }
    fmt.Println(pool.data)
    fmt.Println(pool.head)
    for i := 0; i < 10; i++ {
        pool.Push()
    }
    fmt.Println(pool.data)
    fmt.Println(pool.head)
    */
}

