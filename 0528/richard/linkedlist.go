package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type List struct {
	size int
	head *Node
}

func (l *List) AddAsTail(v int) error {
	// If head is nil, create it first
	if l.IsEmpty() {
		l.head = &Node{value: v}
		return nil
	}

	// Go to the last element of list
	temp := l.head
	counter := 0
	for temp.next != nil {
		temp = temp.next
		counter++
	}

	// Check if number of elements do not exceed the Size of list
	if counter >= l.size {
		return errors.New("List already full!")
	}

	temp.next = &Node{value: v, prev: temp}

	return nil
}

// Adds new element to head of list
func (l *List) AddAsHead(v int) error {
	// If head is nil, create it first
	if l.IsEmpty() {
		l.head = &Node{value: v}
		return nil
	}

	// Go to the last element of list
	temp := l.head
	counter := 0
	for temp.next != nil {
		counter++
		temp = temp.next
	}

	// Check if number of elements do not exceed the Size of list
	if counter >= l.size {
		return errors.New("List already full!")
	}

	newNode := &Node{value: v, prev: temp}
	newNode.next = l.head

	l.head.prev = newNode
	l.head = newNode

	return nil
}

// Check if list is empty
func (l *List) IsEmpty() (r bool) {
	if l.head == nil {
		r = true
	}
	return
}

// Remove element from list
func (l *List) Remove(v int) {
	// If list is empty
	if l.head == nil {
		return
	}

	// Check head first
	if l.head.value == v {
		l.head = l.head.next
	}

	temp := l.head

	// Cross over all elements
	for temp != nil {
		// Check for value
		if temp.value == v {
			// Remove node and replace pointers if not in tail
			if temp.next != nil {
				temp.prev.next = temp.next
				temp.next.prev = temp.prev
				break
			}

			temp.prev.next = nil
		}

		temp = temp.next
	}
}

// Fully clear list
func (l *List) Clear() {
	l.head = nil
}

// List print method
func (l *List) String() string {
	msg := "List current state {Size: %d, IsEmpty: %t, nodes: "
	if l.head == nil {
		return fmt.Sprintf(msg+"[]}", l.size, l.IsEmpty())
	}

	temp := l.head

	// Print head element first
	msg += fmt.Sprintf("[%d", temp.value)

	// Print other nodes
	for temp.next != nil {
		temp = temp.next
		msg += fmt.Sprintf(", %d", temp.value)
	}

	return fmt.Sprintf(msg+"]}", l.size, l.IsEmpty())
}

// List storage
var list *List

func main() {
	// Create list
	list = &List{size: 11}
	fmt.Println(list)

	// Check Add Tail
	fmt.Println("\n==== ADD TO TAIL OF LIST =====")
	for i := 1; i <= 11; i++ {
		if err := list.AddAsTail(i); err != nil {
			fmt.Printf("ERROR: %-v \n", err)
		}
	}
	fmt.Println(list)

	// Check Remove
	fmt.Println("\n==== Remove from list =====")
	for _, v := range []int{5, 1, 7, 11} {
		list.Remove(v)
		fmt.Println(list)
	}

	// Check Add as head
	fmt.Println("\n==== Add as head list =====")
	list.AddAsHead(100)
	fmt.Println(list)

	// Check clear
	fmt.Println("\n==== Clear List =====")
	list.Clear()
	fmt.Println(list)
}
