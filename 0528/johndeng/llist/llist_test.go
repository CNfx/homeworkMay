package llist

import (
	"fmt"
	"testing"
)

func TestLlistAppend(t *testing.T) {
	llist := new(SinglyLinkedList)

	node := Node{value: "hello"}
	llist.Append(&node)

	if llist.Length() != 1 {
		t.Error("Unexpected length")
	}
}

func TestLlistRemove(t *testing.T) {
	llist := new(SinglyLinkedList)

	node := Node{value: "hello"}
	llist.Append(&node)

	llist.Remove(&node)

	if llist.Length() != 0 {
		t.Error("Unexpected length")
	}
}

func TestLlistFind(t *testing.T) {
	llist := new(SinglyLinkedList)

	node := Node{value: "hello"}
	llist.Append(&node)

	n := llist.Find("hello")

	if n == nil {
		t.Error("Unexpect node find")
	}
}

func BenchmarkLlistAppend(b *testing.B) {

	llist := new(SinglyLinkedList)

	for n := 0; n < b.N; n++ {
		node := Node{value: fmt.Sprintf("%d", n)}
		llist.Append(&node)
	}

}

func BenchmarkLlistRemove(b *testing.B) {

	llist := new(SinglyLinkedList)

	for n := 0; n < b.N; n++ {
		node := Node{value: fmt.Sprintf("%d", n)}
		llist.Append(&node)

		llist.Remove(&node)
	}
}

func BenchmarkLlistFind(b *testing.B) {

	llist := new(SinglyLinkedList)

	for n := 0; n < b.N; n++ {
		v := fmt.Sprintf("%d", n)
		node := Node{value: v}
		llist.Append(&node)

		llist.Find(v)
	}

}
