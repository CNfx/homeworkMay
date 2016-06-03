package list

import "errors"

var (
	ErrNextNodeNotExist = errors.New("next node not exist")
	ErrPrevNodeNotExist = errors.New("prev node not exist")
	ErrNodeAlreadyExist = errors.New("node already exist")
	ErrListIsFull       = errors.New("list is full")
)

type List struct {
	prevArr []int
	nextArr []int
	ptrArr  []*Node
	head    int
	tail    int
}

type Node struct {
	idx  int
	name string
	prev int
	next int
	list *List
}

func NewList(cap int) *List {
	return &List{
		prevArr: make([]int, cap),
		nextArr: make([]int, cap),
		ptrArr:  make([]*Node, cap),
	}
}

func (l *List) First() *Node {
	if l.head == 0 {
		return nil
	}

	return l.ptrArr[l.head]
}

func (n *Node) Next() *Node {
	return n.list.ptrArr[n.next]
}

func (n *Node) Prev() *Node {
	return n.list.ptrArr[n.prev]
}

func (l *List) Find(name string) *Node {
	if l.First() == nil {
		return nil
	}

	for node := l.First(); node != nil; node = node.Next() {
		if node.name == name {
			return node
		}
	}

	return nil
}

func (l *List) InsertBefore(nextName, name string) error {
	node := l.Find(name)
	if node != nil {
		return ErrNodeAlreadyExist
	}

	nodeIdx, err := l.findAIdx()
	if err != nil {
		return err
	}
	node = &Node{idx: nodeIdx, name: name, list: l}

	nextNode := l.Find(nextName)
	if nextNode == nil {
		return ErrNextNodeNotExist
	}

	prevIdx := nextNode.prev
	prevNode := l.ptrArr[prevIdx]
	if prevNode == nil {
		return ErrPrevNodeNotExist
	}

	node.next = nextNode.idx
	node.prev = prevNode.idx
	nextNode.prev = node.idx
	prevNode.next = node.idx

	l.nextArr[node.idx] = nextNode.idx
	l.prevArr[node.idx] = prevNode.idx
	l.ptrArr[node.idx] = node

	return nil
}

func (l *List) findAIdx() (int, error) {
	for idx, ptr := range l.ptrArr {
		if ptr == nil {
			return idx, nil
		}
	}

	return 0, ErrListIsFull
}
