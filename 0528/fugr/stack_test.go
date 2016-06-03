package structure

import "testing"

func TestStack(t *testing.T) {
	cap := 10
	s := NewStack(cap)

	_, err := s.Pop()
	if err != ErrStackEmpty {
		t.Error("pop empty stack should with error %s", ErrStackEmpty)
	}

	for i := 0; i < cap; i++ {
		err := s.Push(i)
		if err != nil {
			t.Error(i, err.Error())
		}
	}
	err = s.Push(11)
	if err != ErrStackFull {
		t.Errorf("push data to full stack should with error:%s", ErrStackFull)
	}

	for i := 0; i < cap; i++ {
		dat, err := s.Pop()
		if err != nil {
			t.Error(i, err)
		} else if want := cap - dat - 1; i != want {
			t.Error("want %d got %d", want, dat)
		}

		t.Log(i, dat)
	}

	_, err = s.Pop()
	if err != ErrStackEmpty {
		t.Errorf("pop empty stack should with error:%s", ErrStackEmpty)
	}
}
