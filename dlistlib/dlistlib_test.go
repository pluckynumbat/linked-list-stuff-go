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

