package queue

import (
	"fmt"
	"testing"
)

const dataLen = 10

func TestQueuePush(t *testing.T) {

	q := NewQueue(dataLen)

	for i := dataLen; i > 0; i-- {
		q.Push("Course")
	}

	if q.space == dataLen || q.space != 0 {
		t.Error("Unexpect space cost")
	}
}

func TestQueuePop(t *testing.T) {

	q := NewQueue(dataLen)

	for i := dataLen; i > 0; i-- {
		q.Push(fmt.Sprintf("Course %d", i))
	}

	var data string

	for i := dataLen; i > 0; i-- {
		data, _ = q.Pop()
	}

	if data != "Course 1" {
		t.Error("Unexpect data poped")
	}

}

func BenchmarkQueuePush(b *testing.B) {
	q := NewQueue(b.N)

	for n := 0; n < b.N; n++ {
		q.Push("Course")
	}
}

func BenchmarkQueuePop(b *testing.B) {
	q := NewQueue(b.N)

	for i := b.N; i > 0; i-- {
		q.Push(fmt.Sprintf("Course %d", i))
	}

	for n := 0; n < b.N; n++ {
		q.Pop()
	}
}
