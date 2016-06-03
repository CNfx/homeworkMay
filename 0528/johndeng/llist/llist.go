package llist

type SinglyLinkedList struct {
	front  *Node
	length int
}

type Node struct {
	value string
	next  *Node
}

func (s *SinglyLinkedList) Append(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		currentNode := s.front

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = n
	}

	s.length++
}

// Remove removes node n from list s
func (s *SinglyLinkedList) Remove(n *Node) {
	if s.front == n {
		s.front = n.next
		s.length--
	} else {
		currentNode := s.front

		for currentNode != nil && currentNode.next != nil && currentNode.next != n {
			currentNode = currentNode.next
		}

		if currentNode.next == n {
			currentNode.next = currentNode.next.next
			s.length--
		}
	}
}

func (s *SinglyLinkedList) Find(value interface{}) *Node {
	currentNode := s.front
	for currentNode != nil && currentNode.value != value && currentNode.next != nil {
		currentNode = currentNode.next
	}

	return currentNode
}

func (s *SinglyLinkedList) Length() int {
	return s.length
}
