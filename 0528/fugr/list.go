package structure

import (
	"errors"
	"fmt"
)

var (
	ErrNodeNotFound = errors.New("Node Not Found")
	ErrListEmpty    = errors.New("list empty")
)

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	id   int
	Prev *Node
	Next *Node
}

func NewList() List {
	return List{}
}

func (l *List) Insert(dat int) {
	node := &Node{id: dat}

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	for next := l.head; next != nil; next = next.Next {
		// Increment order
		if next.id < dat {
			if next.Next == nil {
				// insert the next back and node is the tail
				next.Next = node
				l.tail = node
				node.Prev = next
				return
			}
		} else {
			// insert the next front
			node.Prev = next.Prev
			next.Prev = node
			node.Next = next
			if node.Prev == nil {
				l.head = node
			} else {
				node.Prev.Next = node
			}
			return
		}
	}
}

func (l *List) RemoveNode(id int) error {
	if l.head == nil {
		// maybe unnecessary
		return ErrListEmpty
	}

	found := false

	for next := l.head; next != nil; next = next.Next {
		if next.id != id {
			continue
		}

		// remove the node
		switch {
		case next.Prev == nil && next.Next == nil:
			// only one Node
			l.head = nil
			l.tail = nil

		case next.Prev == nil:
			// the first node
			l.head = next.Next
			next.Next.Prev = nil

		case next.Next == nil:
			// next is the tail
			l.tail = next.Prev
			next.Prev.Next = nil

		default:
			next.Prev.Next = next.Next
			next.Next.Prev = next.Prev
		}

		found = true
	}

	if found {
		return nil
	}

	return ErrNodeNotFound
}

func (l *List) Search(id int) (*Node, error) {
	for next := l.head; next != nil; next = next.Next {
		if next.id == id {
			return next, nil
		}
	}

	return nil, ErrNodeNotFound
}

func (l List) String(reverse bool) string {
	if !reverse {
		if l.head == nil {
			return "head -->nil<-- tail"
		}
		s := "head"
		for next := l.head; next != nil; next = next.Next {
			if next.Prev == nil {
				s += fmt.Sprintf(" -->%d", next.id)
			} else {
				s += fmt.Sprintf("<-->%d", next.id)
			}
		}
		return s + "<-- tail"
	}

	// reverse
	if l.tail == nil {
		return "tail -->nil<-- head"
	}
	s := "tail"
	for prev := l.tail; prev != nil; prev = prev.Prev {
		if prev.Next == nil {
			s += fmt.Sprintf(" -->%d", prev.id)
		} else {
			s += fmt.Sprintf("<-->%d", prev.id)
		}
	}

	return s + "<-- head"
}
