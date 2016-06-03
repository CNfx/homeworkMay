package pool

import "testing"

func TestNewPool(t *testing.T) {
	p := NewPool(10)
	if p == nil {
		t.Error("pool must not be nil")
	}
}

func TestPush(t *testing.T) {
	p := NewPool(10)
	p.Pop()

	var data = []struct {
		pushValue   *Object
		expectedErr error
	}{
		{&Object{1}, nil},
		{&Object{2}, ErrPoolFull},
	}

	for _, d := range data {
		err := p.Push(d.pushValue)
		if err != d.expectedErr {
			t.Errorf(
				"push object fail, push value: %v, expectedErr: %v, got: %v",
				d.pushValue,
				d.expectedErr,
				err,
			)
		}
	}
}

func TestPop(t *testing.T) {
	p := NewPool(2)

	var data = []struct {
		expectedErr error
	}{
		{nil},
		{nil},
		{ErrPoolEmpty},
	}

	for _, d := range data {
		_, err := p.Pop()
		if err != d.expectedErr {
			t.Errorf(
				"pop object fail, expectedErr: %v, got: %v",
				d.expectedErr,
				err,
			)
		}
	}
}
