package linkedListLibrary

import (
	"strings"
	"testing"
)

func TestStringForm(t *testing.T) {
	list := ConstructList(ConstructNode("1"))
	want := "1->null"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToEnd(t *testing.T) {
	list := ConstructList(ConstructNode("1"))
	list.AddToEnd("2")
	want := "1->2->null"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToBeginning(t *testing.T) {
	list := ConstructList(ConstructNode("1"))
	list.AddToBeginning("2")
	want := "2->1->null"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestReverse(t *testing.T) {
	list := ConstructList(ConstructNode("2"))
	list.AddToEnd("1")
	list.AddToBeginning("3")
	list.Reverse()
	want := "1->2->3->null"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}
