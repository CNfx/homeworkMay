package structure

import "testing"

func TestQueue(t *testing.T) {
	cap := 10
	q := NewQueue(cap)

	_, err := q.Pop()
	if err != ErrQueueEmpty {
		t.Errorf("pop empty queue should with error:%s", ErrQueueEmpty)
	}

	for i := 0; i < cap; i++ {
		err := q.Push(i)
		if err != nil {
			t.Error(i, err)
		}
	}

	if err := q.Push(11); err != ErrQueueFull {
		t.Errorf("push data to full queue should with error:%s", ErrQueueFull)
	}

	for i := 0; i < cap; i++ {
		dat, err := q.Pop()
		if err != nil {
			t.Error(i, err)
		}
		if i != dat {
			t.Error("want %d got %d", i, dat)
		}
		t.Log(i, dat)
	}

	_, err = q.Pop()
	if err != ErrQueueEmpty {
		t.Errorf("pop empty queue should with error:%s", ErrQueueEmpty)
	}

}

func TestConcurrencyQueue(t *testing.T) {
	q := NewSafeQueue(10)
	ch := make(chan struct{}, 1)

	go func(q *SafeQueue, ch chan struct{}) {
		var (
			dat = 0
			err error
		)
		for i, max := 0, 100+10; i < max; i++ {
			dat, err = q.Pop()

			t.Log(i, "POP ", dat, err)
		}

		ch <- struct{}{}

	}(q, ch)

	go func(q *SafeQueue, ch chan struct{}) {
		for i, max := 0, 100; i < max; i++ {

			t.Log(i, "PUSH", q.Push(i))
		}

		ch <- struct{}{}

	}(q, ch)

	t.Log("1", <-ch)
	t.Log("2", <-ch)
}
