package main

import (
    "errors"
)

var (
    ErrOverFlow = errors.New("range overflow") 
)

type Node struct {
    Data int
    Next *Node
}

type LinkedList struct {
    Head *Node 
    Tail *Node
}

func (l *LinkedList) Append(x int) {
    var node Node
    node.Data = x
    if l.Head == nil {
        l.Head = &node
        l.Tail = &node
    } else {
        l.Tail.Next = &node
        l.Tail = &node
    }
}

func (l *LinkedList) Insert(x, index int) error {
    cur := l.Head
    cur_index := 0
    var node Node
    if index < 0 || index > len(l.Search()) {
        return ErrOverFlow
    } else if index == len(l.Search()) {
        node.Data = x
        l.Tail.Next = &node
        l.Tail = &node
        return nil
    }
    for cur_index < index {
        cur = cur.Next
        cur_index++
    }
    node.Data = x
    node.Next = cur.Next
    cur.Next = &node
    return nil
}


func (l *LinkedList) Search() []Node {
    var s []Node
    cur := l.Head
    for {
        if cur.Next == nil {
            s = append(s, *cur)
            break
        }
        s = append(s, *cur)
        cur = cur.Next
    }
    return s
}


func (l *LinkedList) Remove(index int) {
    cur := l.Head
    cur_index := 0
    for cur_index < (index - 1)  {
        cur = cur.Next
        cur_index++
    }
    cur.Next = cur.Next.Next
}
