package main

import "fmt"

type Node struct {
	Value int
}

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []*Node
	count int
}

func (this *Stack) Push(n *Node) {
	this.nodes = append(this.nodes[:this.count], n)
	this.count++
}

func (this *Stack) Pop() *Node {
	if this.count == 0 {
		return nil
	}
	this.count--
	return this.nodes[this.count]
}

func main() {
	s := NewStack()
	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})
	fmt.Println(s.Pop().Value, s.Pop().Value, s.Pop().Value, s.Pop())
}
