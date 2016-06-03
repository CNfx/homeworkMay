package structure

import "testing"

func TestList(t *testing.T) {
	list := NewList()

	err := list.RemoveNode(1)
	if err != ErrListEmpty {
		t.Error("Unexpected,want %s got %s", ErrListEmpty, err)
	}

	list.Insert(5)
	list.Insert(8)
	list.Insert(11)
	list.Insert(3)
	list.Insert(33)
	list.Insert(12)
	list.Insert(19)
	list.Insert(7)
	list.Insert(30)
	list.Insert(20)

	got := list.String(false)
	if got != "head -->3<-->5<-->7<-->8<-->11<-->12<-->19<-->20<-->30<-->33<-- tail" {
		t.Errorf("Insert error,got:'%s'", got)
	}
	t.Log(got)

	err = list.RemoveNode(5)
	if err != nil {
		t.Error("Remove error:%s", err)
	}

	err = list.RemoveNode(33)
	if err != nil {
		t.Error("Remove error:%s", err)
	}

	err = list.RemoveNode(11)
	if err != nil {
		t.Error("Remove error:%s", err)
	}

	err = list.RemoveNode(10)
	if err != ErrNodeNotFound {
		t.Error("Unexpected,want %s got %s", ErrNodeNotFound, err)
	}

	got = list.String(false)
	if got != "head -->3<-->7<-->8<-->12<-->19<-->20<-->30<-- tail" {
		t.Errorf("Insert error,got:'%s'", got)
	}
	t.Log(got)

	node, err := list.Search(12)
	if err != nil {
		t.Error(err)
	}
	if node.id != 12 {
		t.Error("got wrong node %d", node.id)
	}

	node, err = list.Search(11)
	if err != ErrNodeNotFound {
		t.Error("Unexpected,want %s got %s", ErrNodeNotFound, err)
	}

	got = list.String(true)
	if got != "tail -->30<-->20<-->19<-->12<-->8<-->7<-->3<-- head" {
		t.Errorf("Insert error,got:'%s'", got)
	}
	t.Log(got)
}
