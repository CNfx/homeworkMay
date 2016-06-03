package stack

import (
	"fmt"
	"testing"
)

const dataLen = 10

func TestStackPush(t *testing.T) {

	s := NewStack(dataLen)

	for i := dataLen; i > 0; i-- {
		s.Push("Course")
	}

	if s.space == dataLen || s.space != 0 {
		t.Error("Unexpect space cost")
	}
}

func TestStackPop(t *testing.T) {

	s := NewStack(dataLen)

	for i := dataLen; i > 0; i-- {
		s.Push(fmt.Sprintf("Course %d", i))
	}

	var data string

	for i := dataLen; i > 0; i-- {
		data, _ = s.Pop()
	}

	if data != "Course 10" {
		t.Error("Unexpect data poped")
	}

}

func BenchmarkStackPush(b *testing.B) {
	s := NewStack(b.N)

	for n := 0; n < b.N; n++ {
		s.Push("Course")
	}
}

func BenchmarkStackPop(b *testing.B) {
	s := NewStack(b.N)

	for i := b.N; i > 0; i-- {
		s.Push(fmt.Sprintf("Course %d", i))
	}

	for n := 0; n < b.N; n++ {
		s.Pop()
	}
}
