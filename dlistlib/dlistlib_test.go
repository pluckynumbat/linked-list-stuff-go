package dlistlib

import (
	"testing"
)

func TestNilNodeString(t *testing.T) {
	var node *Node
	got := node.String()
	want := "nil"

	if got != want {
		t.Errorf("Incorrect string for a nil node: want %v, got %v", want, got)
	}
}

func TestNodeString(t *testing.T) {
	node := &Node{nil, "a", nil}
	got := node.String()
	want := "<-a->"

	if got != want {
		t.Errorf("Incorrect string for a non-nil node: want %v, got %v", want, got)
	}
}

func TestNilListString(t *testing.T) {
	var dlist *DoublyLinkedList

	got := dlist.String()
	want := "nil"

	if got != want {
		t.Errorf("Incorrect string for a nil list: want %v, got %v", want, got)
	}
}

func TestSingleElementListString(t *testing.T) {
	node := &Node{nil, "a", nil}
	dlist := &DoublyLinkedList{node, node}
	got := dlist.String()
	want := "nil<-a->nil"

	if got != want {
		t.Errorf("Incorrect string for a single element list: want %v, got %v", want, got)
	}
}

func TestListString(t *testing.T) {
	var head, tail *Node
	head = &Node{nil, "a", nil}
	tail = &Node{nil, "b", nil}

	head.next = tail
	tail.prev = head

	dlist := &DoublyLinkedList{head, tail}
	got := dlist.String()
	want := "nil<-a<=>b->nil"

	if got != want {
		t.Errorf("Incorrect string for a single element list: want %v, got %v", want, got)
	}
}

func TestIsNil(t *testing.T) {
	var dlist *DoublyLinkedList

	want := true
	got := dlist.IsNil()

	if got != want {
		t.Errorf("Incorrect value of IsNil on nil list: want %v, got %v", want, got)
	}

	dlist = &DoublyLinkedList{}

	want = false
	got = dlist.IsNil()

	if got != want {
		t.Errorf("Incorrect value of IsNil on a non nil list: want %v, got %v", want, got)
	}

	dlist = new(DoublyLinkedList)

	want = false
	got = dlist.IsNil()

	if got != want {
		t.Errorf("Incorrect value of IsNil on a non nil list: want %v, got %v", want, got)
	}
}

