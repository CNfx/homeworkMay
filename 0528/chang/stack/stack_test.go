package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack(10)
	if s == nil {
		t.Error("stack must not be nil")
	}
}

func TestPush(t *testing.T) {
	stackObj := NewStack(2)

	var data = []struct {
		pushValue   int
		expectedErr error
	}{
		{1, nil},
		{2, nil},
		{3, ErrStackFull},
	}

	for _, d := range data {
		err := stackObj.Push(d.pushValue)
		if err != d.expectedErr {
			t.Errorf(
				"test push full fail: pushed %d, expected: %v, got: %v",
				d.pushValue,
				d.expectedErr,
				err,
			)
		}
	}
}

func TestPop(t *testing.T) {
	stackObj := NewStack(10)

	_, err := stackObj.Pop()
	if err != ErrStackEmpty {
		t.Errorf("pop empty stack should return stack empty error")
	}

	stackObj.Push(1)
	stackObj.Push(2)

	stackObj.Pop()
	stackObj.Pop()
	_, err = stackObj.Pop()
	if err != ErrStackEmpty {
		t.Errorf("pop empty stack should return stack empty error")
	}
}
