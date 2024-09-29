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
		t.Errorf("GetStringForm Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToEnd(t *testing.T) {
	list := constructList(constructNode("1"))
	list.AddToEnd("2")
	want := "1->2->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("AddToEnd Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestAddToBeginning(t *testing.T) {
	list := constructList(constructNode("1"))
	list.AddToBeginning("2")
	want := "2->1->nil"
	have := list.GetStringForm()
	if strings.Compare(want, have) != 0 {
		t.Errorf("AddToBeginning Fails: wanted '%s', got '%s'", want, have)
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
		t.Errorf("Reverse Fails: wanted '%s', got '%s'", want, have)
	}
}

func TestConstructFromValues(t *testing.T) {
	list1 := constructList(constructNode("t"))
	list1.AddToBeginning("a")
	list1.AddToEnd("e")
	list1.AddToBeginning("w")
	list1.AddToEnd("r")

	list2 := ConstructFromValues("w", "a", "t", "e", "r")

	want := list1.GetStringForm()
	have := list2.GetStringForm()

	if strings.Compare(want, have) != 0 {
		t.Errorf("ConstructFromValues Fails: wanted '%s', got '%s'", want, have)
	}
}
