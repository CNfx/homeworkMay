package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue(10)
	if q == nil {
		t.Error("queue must not be nil")
	}
}

func TestPush(t *testing.T) {
	q := NewQueue(2)

	var data = []struct {
		pushValue   int
		expectedErr error
	}{
		{1, nil},
		{2, nil},
		{3, ErrQueueFull},
	}

	for _, d := range data {
		err := q.Push(d.pushValue)
		if err != d.expectedErr {
			t.Errorf(
				"queue push fail: expected: %v, got: %v",
				d.expectedErr,
				err,
			)
		}
	}
}

func TestPop(t *testing.T) {
	q := NewQueue(2)
	q.Push(1)
	q.Push(2)

	var data = []struct {
		expectedValue int
		expectedErr   error
	}{
		{1, nil},
		{2, nil},
		{0, ErrQueueEmpty},
	}

	for _, d := range data {
		v, err := q.Pop()
		if v != d.expectedValue || err != d.expectedErr {
			t.Errorf(
				"queue pop fail: expected value: %d, got value: %d,  expectedErr: %v, got: %v",
				d.expectedValue,
				v,
				d.expectedErr,
				err,
			)
		}
	}
}
