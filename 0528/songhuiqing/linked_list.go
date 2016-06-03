package main

import (
	"errors"
)

type List struct {
	Length int //长度
	Head   *Node //头
	Tail   *Node //尾
}

func NewList() *List {
	l := new(List)
	l.Length = 0
	return l
}

type Node struct {
	Value int
	Prev  *Node //前一个
	Next  *Node //下一个
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (l *List) Len() int {
	return l.Length
}

func (l *List) IsEmpty() bool {
	return l.Length == 0
}

/**
*	加入作为第一个子节点
**/
func (l *List) Prepend(value int) {
	node := NewNode(value)
	if l.Len() == 0 {//如果list为空，则让新创建的NODE为HEAD，同时为TAIL
		l.Head = node
		l.Tail = l.Head
	} else {//如果不为空
		formerHead := l.Head//把原有的头节点取出保存在formerHead
		formerHead.Prev = node//设置formerHead的

		node.Next = formerHead
		l.Head = node
	}

	l.Length++
}

/**
 *	加入作为最后一个子节点
 **/
func (l *List) Append(value int) {
	node := NewNode(value)

	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerTail := l.Tail
		formerTail.Next = node

		node.Prev = formerTail
		l.Tail = node
	}

	l.Length++
}

func (l *List) Add(value int, index int) error {
	if index > l.Len() {
		return errors.New("Index out of range")
	}

	node := NewNode(value)

	if l.Len() == 0 || index == 0 {
		l.Prepend(value)
		return nil
	}

	if l.Len()-1 == index {
		l.Append(value)
		return nil
	}

	nextNode, _ := l.Get(index)
	prevNode := nextNode.Prev

	prevNode.Next = node
	node.Prev = prevNode

	nextNode.Prev = node
	node.Next = nextNode

	l.Length++

	return nil
}

func (l *List) Get(index int) (*Node, error) {
	if index > l.Len() {
		return nil, errors.New("Index out of range")
	}

	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node, nil
}

/*func (l *List) Remove(value int) error {
	if l.Len() == 0 {
		return errors.New("Empty list")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	found := 0
	for n := l.Head; n != nil; n = n.Next {

		if *n.Value.(*Node) == value && found == 0 {
			n.Next.Prev, n.Prev.Next = n.Prev, n.Next
			l.Length--
			found++
		}
	}

	if found == 0 {
		return errors.New("Node not found")
	}

	return nil
}*/

func main() {
	list := NewList()
	println(list.Len())
	list.Prepend(888)
	println(list.Len())
	list.Get(0)
}
