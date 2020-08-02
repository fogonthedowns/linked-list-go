package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	ll := LinkedList{}
	ll.Push(&Node{value: 1})
	ll.Push(&Node{value: 2})
	ll.Push(&Node{value: 3})
	node := ll.Pop(2)
	if node.value != 3 {
		t.Errorf("got %d, want %d", node.value, 3)
	}

	if ll.length != 2 {
		t.Errorf("got %d, want %d", ll.length, 2)
	}

	node = ll.Pop(2)
	if node != nil {
		t.Errorf("got %v, want %v", node, nil)
	}

	if ll.length != 2 {
		t.Errorf("got %d, want %d", ll.length, 2)
	}

	ll.Delete(0)
	if ll.length != 1 {
		t.Errorf("got %d, want %d", ll.length, 1)
	}

	if ll.start.value != 1 && ll.end.value != 1 {
		t.Errorf("got %d and got %d , want %d and want %d", ll.start.value, ll.end.value, 1, 1)
	}
	ll.Push(&Node{value: 31})
	if ll.start.value != 1 && ll.end.value != 31 {
		t.Errorf("got %d and got %d , want %d and want %d", ll.start.value, ll.end.value, 1, 31)
	}
	if ll.length != 2 {

		t.Errorf("got %d, want %d", ll.length, 2)
	}
}
