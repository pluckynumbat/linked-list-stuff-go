package linkedListLibrary

import (
	"strings"
	"testing"
)

func TestStringForm(t *testing.T) {
	list := constructList(constructNode("1"))
	want := "1->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToEnd(t *testing.T) {
	list := constructList(constructNode("1"))
	list.AddToEnd("2")
	want := "1->2->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToBeginning(t *testing.T) {
	list := constructList(constructNode("1"))
	list.AddToBeginning("2")
	want := "2->1->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestReverse(t *testing.T) {
	list := constructList(constructNode("2"))
	list.AddToEnd("1")
	list.AddToBeginning("3")
	list.Reverse()
	want := "1->2->3->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Fatalf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}
