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

